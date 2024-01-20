package main

import (
	"io"
	"net/http"

	"github.com/rs/cors"
	"github.com/uptrace/bunrouter"
	"lure.sh/lure/pkg/loggerctx"
	"lure.sh/lure/pkg/search"
	"go.elara.ws/logger/log"
)

func registerAPI(mux *bunrouter.Router) {
	g := mux.NewGroup(
		"/api",
		bunrouter.WithMiddleware(apiErrorHandler, corsMiddleware()),
	)
	
	g.GET("/*path", func(w http.ResponseWriter, req bunrouter.Request) error {
		return HTTPError{404, "endpoint not found: " + req.URL.Path}
	})

	g.GET("/search", func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := loggerctx.With(req.Context(), log.Logger)
		pkgs, err := search.Search(ctx, searchOptions(req.Request))
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, translatePkgs(getLanguages(req.Request), pkgs))
	})

	g.GET("/pkg/:repo/:pkg", func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := loggerctx.With(req.Context(), log.Logger)
		pkg, err := search.GetPkg(ctx, req.Param("repo"), req.Param("pkg"))
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, translatePkg(getLanguages(req.Request), pkg))
	})

	g.GET("/pkg/:repo/:pkg/script", func(w http.ResponseWriter, req bunrouter.Request) error {
		ctx := loggerctx.With(req.Context(), log.Logger)		
		script, err := search.GetScript(ctx, req.Param("repo"), req.Param("pkg"))
		if err != nil {
			return err
		}
		defer script.Close()

		if req.URL.Query().Get("highlight") == "true" {
			data, err := highlight(script)
			if err != nil {
				return err
			}
			w.Header().Set("Content-Type", "text/html")
			_, err = io.WriteString(w, string(data))
			return err
		} else {
			w.Header().Set("Content-Type", "text/x-shellscript")
			_, err = io.Copy(w, script)
			return err
		}
	})
}

func corsMiddleware() bunrouter.MiddlewareFunc {
	ch := cors.AllowAll()
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return bunrouter.HTTPHandler(ch.Handler(next))
	}
}
