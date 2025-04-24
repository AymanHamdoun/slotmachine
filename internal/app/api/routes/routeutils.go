package routes

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go-backend/internal/app/api/middleware"
	"net/http"
)

type handler[Req any] interface {
	Serve(context.Context, Req, http.ResponseWriter, *http.Request)
}

func getRequest[Req any](ctx context.Context) *Req {
	val := ctx.Value(middleware.BindingRequestCtxKey{})
	request, ok := val.(*Req)
	if ok {
		return request
	}
	return nil
}

func registerGET[Req any](
	ctx context.Context,
	router *chi.Mux,
	paths []string,
	handler handler[Req]) {
	for _, path := range paths {
		middlewares := []func(next http.Handler) http.Handler{
			middleware.BindingMiddleware[Req](),
		}
		router.With(middlewares...).Get(path, genericHandler(ctx, path, handler))
	}
}

func registerPOST[Req any](
	ctx context.Context,
	router *chi.Mux,
	paths []string,
	handler handler[Req]) {
	for _, path := range paths {
		middlewares := []func(next http.Handler) http.Handler{
			middleware.BindingMiddleware[Req](),
		}
		router.With(middlewares...).Post(path, genericHandler(ctx, path, handler))
	}
}

func genericHandler[Req any](_ context.Context, _ string, handler handler[Req]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := getRequest[Req](r.Context())
		handler.Serve(r.Context(), *request, w, r)
	}
}
