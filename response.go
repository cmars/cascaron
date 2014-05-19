package cascaron

import (
	"net/http"
)

func notImplemented(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotImplemented)
}
