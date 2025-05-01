package viewhandlers

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"net/http"
)

func _view(ctx context.Context, w http.ResponseWriter, r *http.Request, viewPath string, viewData any) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, fmt.Sprintf("web/dist/%s", viewPath))
}

func _templ(w http.ResponseWriter, r *http.Request, component templ.Component) {
	err := component.Render(r.Context(), w)
	if err != nil {
		return
	}
}
