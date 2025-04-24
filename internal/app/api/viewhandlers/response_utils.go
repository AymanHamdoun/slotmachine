package viewhandlers

import (
	"context"
	"fmt"
	"net/http"
)

func _view(ctx context.Context, w http.ResponseWriter, r *http.Request, viewPath string, viewData any) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, fmt.Sprintf("web/dist/%s", viewPath))
}
