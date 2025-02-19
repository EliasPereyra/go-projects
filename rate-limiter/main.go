package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"golang.org/x/time/rate"
)

type Response struct {
	Message string `json:"message"`
}

func getIp(r *http.Request) string {
	// por medio del struct del Request obtengo la direccion de su IP
	// Ahora, en el request me dice que devuelve el IP + PORT, por lo que
	// necesito extraer solamente la IP y lo devuelvo
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Error when parsing the IP %v", err)
		return ""
	}
	return host
}

func rateLimitierMiddleware(next http.Handler, limit rate.Limit, burst int) http.Handler {
	// Para un rate limiter, necesito ver si la IP está en un bucket del limiter
	// Por ello genero un estructura para las ips y revisar que estan siendo limitadas
	// o sino asegurar que lo estén
	ipLimiterBucket := make(map[string]*rate.Limiter)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtengo la IP
		ip := getIp(r)

		// Ahora creo un limiter para la IP, si no lo tiene todavia
		limiter, exists := ipLimiterBucket[ip]
		if !exists {
			limiter = rate.NewLimiter(limit, burst)
			ipLimiterBucket[ip] = limiter
		}

		if !limiter.Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{"error": "Too many requests"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello there!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greetHandler)

	handler := rateLimitierMiddleware(mux, rate.Limit(1), 10)

	log.Println("Server started at port 8080")
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		log.Fatal(err)
	}
}
