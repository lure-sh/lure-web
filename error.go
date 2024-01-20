package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/uptrace/bunrouter"
	"go.elara.ws/logger/log"
	"go.elara.ws/salix"
)

type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"error"`
}

func (he HTTPError) Error() string {
	return he.Msg
}

func (he HTTPError) StatusText() string {
	return http.StatusText(he.StatusCode)
}

func httpError(err error) *HTTPError {
	switch err := err.(type) {
	case nil:
		return nil
	case HTTPError:
		return &err
	case *HTTPError:
		return err
	default:
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return &HTTPError{404, err.Error()}
		default:
			return &HTTPError{500, err.Error()}
		}
	}
}

func siteErrorHandler(ns *salix.Namespace) bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		tmpl, ok := ns.GetTemplate("error.html")
		if !ok {
			return basicErrorHandler(next)
		}


		return func(w http.ResponseWriter, req bunrouter.Request) error {
			err := next(w, req)
			if err != nil {
				he := httpError(err)
				w.WriteHeader(he.StatusCode)
				err := tmpl.
					WithVarMap(map[string]any{"error": he}).
					Execute(w)
				if err != nil {
					log.Error("Error while executing error template").Err(err).Send()
				}
			}
			return err
		}
	}
}

func apiErrorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)
		if err != nil {
			he := httpError(err)
			w.WriteHeader(he.StatusCode)
			_ = bunrouter.JSON(w, map[string]string{"error": he.Msg})
		}
		return err
	}
}

func basicErrorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)
		if err != nil {
			he := httpError(err)
			http.Error(w, he.Msg, he.StatusCode)
		}
		return err
	}
}
