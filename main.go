package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"github.com/uptrace/bunrouter"
	"go.elara.ws/logger"
	"go.elara.ws/logger/log"
	"lure.sh/lure/pkg/loggerctx"
)

//go:embed static
var static embed.FS

func init() {
	log.Logger = logger.NewPretty(os.Stderr)
}

func main() {
	mux := bunrouter.New()

	fileServer := http.FileServer(http.FS(static))
	mux.Use(cache).GET("/static/*path", bunrouter.HTTPHandler(fileServer))

	ctx := loggerctx.With(context.Background(), log.Logger)

	registerBadge(mux)
	registerSite(mux)
	registerAPI(mux)
	registerWebhook(ctx, mux)

	http.ListenAndServe(":8080", mux)
}

func cache(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return bunrouter.HandlerFunc(func(w http.ResponseWriter, req bunrouter.Request) error {
		w.Header().Set("Cache-Control", "max-age=86400, stale-while-revalidate=86400")
		return next(w, req)
	})
}
