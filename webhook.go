package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/uptrace/bunrouter"
	"go.elara.ws/logger/log"
	"lure.sh/lure/pkg/repos"
)

func registerWebhook(ctx context.Context, mux *bunrouter.Router) {
	g := mux.WithMiddleware(apiErrorHandler)

	g.POST("/webhook", func(w http.ResponseWriter, req bunrouter.Request) error {
		if req.Header.Get("X-GitHub-Event") != "push" {
			return HTTPError{http.StatusBadRequest, "only push events are accepted by this bot"}
		}

		err := verifySecure(req.Request)
		if err != nil {
			return err
		}

		go pullRepos(ctx)
		return nil
	})
}

func verifySecure(req *http.Request) error {
	sigStr := req.Header.Get("X-Hub-Signature-256")
	sig, err := hex.DecodeString(strings.TrimPrefix(sigStr, "sha256="))
	if err != nil {
		return HTTPError{http.StatusBadRequest, "invalid hmac value"}
	}

	secretStr, ok := os.LookupEnv("LURE_WEB_GITHUB_SECRET")
	if !ok {
		return HTTPError{http.StatusInternalServerError, "LURE_WEB_GITHUB_SECRET must be set to the secret used for setting up the github webhook"}
	}
	secret := []byte(secretStr)

	h := hmac.New(sha256.New, secret)
	_, err = io.Copy(h, req.Body)
	if err != nil {
		return err
	}

	if !hmac.Equal(h.Sum(nil), sig) {
		log.Warn("Insecure webhook request").
			Str("from", req.RemoteAddr).
			Bytes("sig", sig).
			Bytes("hmac", h.Sum(nil)).
			Send()

		return HTTPError{http.StatusUnauthorized, "insecure webhook request"}
	}

	return nil
}

var pullMtx sync.Mutex

func pullRepos(ctx context.Context) {
	pullMtx.Lock()
	defer pullMtx.Unlock()

	err := repos.Pull(ctx, nil)
	if err != nil {
		log.Warn("Error while pulling repositories").Err(err).Send()
	}
}
