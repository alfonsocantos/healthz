package basic

import "net/http"
import "fmt"

type health struct {
	Status string `json:"status"`
}

// Healthz this is the Healthz endpoint
func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requesting ....", i)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
