package http

import (
	endpoint "auth/pkg/endpoint"
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeAuthHandler creates the handler logic
func makeAuthHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/auth", http1.NewServer(endpoints.AuthEndpoint, decodeAuthRequest, encodeAuthResponse, options...))
}

// decodeAuthRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AuthRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAuthResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAuthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangePasswordHandler creates the handler logic
func makeChangePasswordHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/change-password", http1.NewServer(endpoints.ChangePasswordEndpoint, decodeChangePasswordRequest, encodeChangePasswordResponse, options...))
}

// decodeChangePasswordRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangePasswordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ChangePasswordRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangePasswordResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangePasswordResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSignUpHandler creates the handler logic
func makeSignUpHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/sign-up", http1.NewServer(endpoints.SignUpEndpoint, decodeSignUpRequest, encodeSignUpResponse, options...))
}

// decodeSignUpRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSignUpResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSignUpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete", http1.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...))
}

// decodeDeleteRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
