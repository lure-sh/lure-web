package main

import (
	"context"
	_ "embed"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"

	"github.com/uptrace/bunrouter"
	"lure.sh/lure/pkg/search"
)

type pkgCtxKey struct{}

//go:embed static/icons/badge-logo.base64
var logoData string

// registerBadge registers the route for generating LURE badges
func registerBadge(mux *bunrouter.Router) {
	// rp is a reverse proxy that will route user requests to shields.io
	// and return the responses, stripped of any irrelevant headers.
	rp := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pkg := pr.In.Context().Value(pkgCtxKey{}).(search.Package)
			pr.Out.URL = genBadgeURL(pkg, pr.In.URL.Query().Get("style"))
			pr.Out.Host = pr.Out.URL.Host
		},
		ModifyResponse: func(r *http.Response) error {
			// Remove all the irrelevant headers from the response.
			r.Header.Del("Alt-Svc")
			r.Header.Del("Cache-Control")
			r.Header.Del("Report-To")
			r.Header.Del("Cf-Cache-Status")
			r.Header.Del("Cf-Ray")
			r.Header.Del("Fly-Request-Id")
			r.Header.Del("Nel")
			r.Header.Del("Report-To")
			r.Header.Del("Server")
			r.Header.Del("Via")
			r.Header.Del("Set-Cookie")
			return nil
		},
	}

	mux.GET("/pkg/:repo/:pkg/badge.svg", func(w http.ResponseWriter, req bunrouter.Request) error {
		repo := req.Param("repo")
		name := req.Param("pkg")

		pkg, err := search.GetPkg(req.Context(), repo, name)
		if err != nil {
			return err
		}

		ctx := context.WithValue(req.Context(), pkgCtxKey{}, pkg)
		rp.ServeHTTP(w, req.Request.WithContext(ctx))
		return nil
	})
}

func genBadgeURL(pkg search.Package, style string) *url.URL {
	v := url.Values{}
	v.Set("label", pkg.Name)
	v.Set("message", genBadgeVersion(pkg))
	v.Set("logo", logoData)
	v.Set("color", "0060A8")
	if style != "" {
		v.Set("style", style)
	}
	return &url.URL{Scheme: "https", Host: "img.shields.io", Path: "/static/v1", RawQuery: v.Encode()}
}

func genBadgeVersion(pkg search.Package) string {
	sb := strings.Builder{}
	if pkg.Epoch != 0 {
		sb.WriteString(strconv.Itoa(int(pkg.Epoch)))
		sb.WriteByte(':')
	}

	sb.WriteString(pkg.Version)

	if pkg.Release != 0 {
		sb.WriteByte('-')
		sb.WriteString(strconv.Itoa(pkg.Release))
	}
	return sb.String()
}
