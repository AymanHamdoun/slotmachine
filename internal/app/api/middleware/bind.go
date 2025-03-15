package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
)

const contentTypeHeader = "Content-Type"

// Content types.
const (
	JSONContentType     = "application/json"
	FormBodyContentType = "application/x-www-form-urlencoded"
)

type bodyCtxKey struct{}

// Content-Type header can include tags like charset and lang which need to be filtered out while binding.
func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func BindingMiddleware[Req any]() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			toBind := new(Req)

			err := bindDataBasedOnContentType(r, toBind)
			if err != nil {
				//bodyBytes, _ := io.ReadAll(r.Body)
				//logger.Ctx(r.Context()).Warn("binding middleware: failed to bind data",
				//	logger.Error(err),
				//	logger.String("path", path),
				//	logger.String("url", r.URL.String()),
				//	logger.String("method", r.Method),
				//	logger.String("payload", string(bodyBytes)),
				//	logger.Any("headers", r.Header),
				//)

				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}

			ctx := context.WithValue(r.Context(), bodyCtxKey{}, toBind)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func bindDataBasedOnContentType[Req any](r *http.Request, toBind *Req) error {
	contentType := filterFlags(r.Header.Get(contentTypeHeader))

	// We have requests that are setting the content-type to json or binary even when the parameters are meant to be read from the query string.
	if r.Method == http.MethodGet {
		contentType = FormBodyContentType
	}

	switch contentType {
	case FormBodyContentType:
		return bindDataFromFormRequest(r, toBind)
	case JSONContentType:
		return bindDataFromJSONRequest(r, toBind)
	default:
		return bindDataFromFormRequest(r, toBind)
	}
}

func bindDataFromFormRequest[Req any](r *http.Request, toBind *Req) error {
	contentType := filterFlags(r.Header.Get(contentTypeHeader))
	// Copy the body
	bodyCopy, _ := io.ReadAll(r.Body)

	// Refill so we can properly parse it
	r.Body = io.NopCloser(bytes.NewReader(bodyCopy))

	// Consume the body
	if err := r.ParseForm(); err != nil {
		return err
	}

	err := binding.Default(r.Method, contentType).Bind(r, toBind)

	// Refill the body case someone needs down the chain
	r.Body = io.NopCloser(bytes.NewReader(bodyCopy))

	if err != nil {
		return fmt.Errorf("failed to decode form request: %w", err)
	}
	return nil
}

func bindDataFromJSONRequest[Req any](r *http.Request, toBind *Req) error {
	bodyBytes, err := io.ReadAll(r.Body)

	// Refill it in case someone needs down the chain
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("could not read request body: %w", err)
	}

	if len(bodyBytes) == 0 {
		return nil
	}

	err = json.Unmarshal(bodyBytes, &toBind)

	if err != nil {
		return fmt.Errorf("could not unmarshal request JSON: %w", err)
	}
	return nil
}
