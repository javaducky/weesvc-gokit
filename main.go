package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc GreetingService
	svc = greetingService{}
	svc = loggingMiddleware{logger, svc}

	greetingHandler := httptransport.NewServer(
		makeGreetingEndpoint(svc),
		decodeGreetingRequest,
		encodeGreetingResponse,
	)

	http.Handle("/greeting", greetingHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
