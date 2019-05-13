package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeGreetingEndpoint(svc GreetingService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(greetingRequest)
		v, err := svc.Greeting(req.Name)
		if err != nil {
			return greetingResponse{v, err.Error()}, nil
		}
		return greetingResponse{v, ""}, nil
	}
}

func decodeGreetingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request greetingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeGreetingResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type greetingRequest struct {
	Name string `json:"name"`
}

type greetingResponse struct {
	Name string `json:"name"`
	Err  string `json:"err,omitempty"`
}
