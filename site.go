package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/uptrace/bunrouter"
	"go.elara.ws/logger/log"
	"go.elara.ws/salix"
	"lure.sh/lure/pkg/loggerctx"
	"lure.sh/lure/pkg/search"
)

//go:embed tmpls
var tmpls embed.FS

func registerSite(mux *bunrouter.Router) {
	tmplFS, err := fs.Sub(tmpls, "tmpls")
	if err != nil {
		panic(err)
	}

	ns := salix.New().
		WithEscapeHTML(true).
		WithWriteOnSuccess(true).
		WithVarMap(map[string]any{
			"now": time.Now,
		})
	err = ns.ParseFSGlob(tmplFS, "*.html")
	if err != nil {
		panic(err)
	}

	g := mux.WithMiddleware(siteErrorHandler(ns))

	g.GET("/*path", func(w http.ResponseWriter, req bunrouter.Request) error {
		if req.Param("path") == "" {
			return ns.ExecuteTemplate(w, "home.html", nil)
		}
		return HTTPError{404, "page not found: " + req.URL.Path}
	})

	g.GET("/about", func(w http.ResponseWriter, req bunrouter.Request) error {
		return ns.ExecuteTemplate(w, "about.html", nil)
	})

	g.GET("/icons", func(w http.ResponseWriter, req bunrouter.Request) error {
		return ns.ExecuteTemplate(w, "icons.html", nil)
	})

	g.GET("/pkgs", func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := loggerctx.With(req.Context(), log.Logger)
		pkgs, err := search.Search(ctx, searchOptions(req.Request))
		if err != nil {
			return err
		}

		return ns.ExecuteTemplate(w, "pkgs.html", map[string]any{
			"pkgs":  translatePkgs(getLanguages(req.Request), pkgs),
			"query": req.URL.Query(),
		})
	})

	g.GET("/pkg/:repo/:pkg", func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := loggerctx.With(req.Context(), log.Logger)
		pkg, err := search.GetPkg(ctx, req.Param("repo"), req.Param("pkg"))
		if err != nil {
			return err
		}

		return ns.ExecuteTemplate(w, "pkg.html", map[string]any{
			"url": getURL(req.Request),
			"pkg": translatePkg(getLanguages(req.Request), pkg),
		})
	})

	g.GET("/pkg/:repo/:pkg/script", func(w http.ResponseWriter, req bunrouter.Request) error {
		repoName := req.Param("repo")
		pkgName := req.Param("pkg")

		ctx := loggerctx.With(req.Context(), log.Logger)
		script, err := search.GetScript(ctx, repoName, pkgName)
		if err != nil {
			return err
		}
		defer script.Close()

		data, err := highlight(script)
		if err != nil {
			return err
		}

		return ns.ExecuteTemplate(w, "script.html", map[string]any{
			"repoName": repoName,
			"pkgName":  pkgName,
			"script":   data,
		})
	})
}

// searchOptions gets a search.Options struct from the request's
// query parameters.
func searchOptions(r *http.Request) search.Options {
	sort := search.SortBy(search.SortByNone)
	switch r.URL.Query().Get("sort") {
	case "name":
		sort = search.SortByName
	case "version":
		sort = search.SortByVersion
	case "repo":
		sort = search.SortByRepo
	}

	filter := search.FilterNone
	switch r.URL.Query().Get("filter") {
	case "inrepo":
		filter = search.FilterInRepo
	case "arch":
		filter = search.FilterSupportsArch
	}

	return search.Options{
		Filter:      filter,
		FilterValue: r.URL.Query().Get("fv"),
		SortBy:      sort,
		Query:       r.URL.Query().Get("q"),
	}
}

// highlight returns an HTML string with syntax highlighting
// for the contents of the given reader.
func highlight(r io.Reader) (salix.HTML, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	lexer := lexers.Get("bash")
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get("modus-vivendi")
	formatter := html.New()
	iterator, err := lexer.Tokenise(nil, string(data))
	if err != nil {
		return "", err
	}

	out := &strings.Builder{}
	err = formatter.Format(out, style, iterator)
	if err != nil {
		return "", err
	}

	return salix.HTML(out.String()), nil
}

// getURL attempts to get the current URL of the request
func getURL(r *http.Request) string {
	u := *r.URL

	xfp := r.Header.Get("X-Forwarded-Proto")
	if xfp == "" {
		if r.TLS == nil {
			u.Scheme = "http"
		} else {
			u.Scheme = "https"
		}
	} else {
		u.Scheme = xfp
	}

	xfh := r.Header.Get("X-Forwarded-Host")
	if xfh == "" {
		if r.URL.Host == "" {
			u.Host = r.Host
		}
	} else {
		u.Host = xfh
	}

	return u.String()
}
