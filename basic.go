package basic

import "net/http"

type health struct {
	Status string `json:"status"`
}

// Healthz this is the Healthz endpoint
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
