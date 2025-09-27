package viewhandlers

// these funcs should be moved to responseutils/response.go
import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
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
