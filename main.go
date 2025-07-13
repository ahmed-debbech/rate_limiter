package main

import (
	"fmt"
	"log"
	"net/http"

	rlHttp "github.com/ahmed-debbech/rate_limiter/http"
	"github.com/ahmed-debbech/rate_limiter/logic"
)

func main() {
	log.Println("Rate Limiter Starting...")
	http.HandleFunc("/", rlHttp.PassHandler)
	fmt.Println("Server listening on :1700")

	go logic.RefreshTokens()
	log.Fatal(http.ListenAndServe(":1700", nil))
}
