package http

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ahmed-debbech/rate_limiter/logic"
)

func PassHandler(w http.ResponseWriter, r *http.Request) {
	// Clone original request
	go func() {
		req, err := http.NewRequest(r.Method, "http://debbech.com"+r.RequestURI, r.Body)
		if err != nil {
			log.Println("ERROR ", err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Copy headers
		req.Header = r.Header.Clone()

		// Append client IP to X-Forwarded-For
		originalIP := r.RemoteAddr
		if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
			req.Header.Set("X-Forwarded-For", ip+", "+strings.Split(originalIP, ":")[0])
		} else {
			req.Header.Set("X-Forwarded-For", strings.Split(originalIP, ":")[0])
		}

		// integrate logic
		err = logic.ConsumeToken(originalIP)
		if err != nil {
			log.Println("ERROR ", err.Error())
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}

		// Forward to internal service
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("ERROR", err.Error())
			http.Error(w, "Could not reach backend servers", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Copy status code and headers
		for key, values := range resp.Header {
			for _, v := range values {
				w.Header().Add(key, v)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}()
}
