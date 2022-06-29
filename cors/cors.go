package cors

import (
	"fmt"
	"net/http"
)

func SetupCORS(w *http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside CORS")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
