package endpoint

import (
	model "auth/internal/model"
	service "auth/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
)

// AuthRequest collects the request parameters for the Auth method.
type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthResponse collects the response parameters for the Auth method.
type AuthResponse struct {
	Jwt string `json:"jwt"`
	Err error  `json:"err"`
}

// MakeAuthEndpoint returns an endpoint that invokes Auth on the service.
func MakeAuthEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthRequest)
		jwt, err := s.Auth(ctx, req.Login, req.Password)
		return &AuthResponse{
			Err: err,
			Jwt: jwt,
		}, nil
	}
}

// Failed implements Failer.
func (r AuthResponse) Failed() error {
	return r.Err
}

// ChangePasswordRequest collects the request parameters for the ChangePassword method.
type ChangePasswordRequest struct {
	Login string `json:"login"`
	Curr  string `json:"curr"`
	New   string `json:"new"`
}

// ChangePasswordResponse collects the response parameters for the ChangePassword method.
type ChangePasswordResponse struct {
	Success bool  `json:"success"`
	Err     error `json:"err"`
}

// MakeChangePasswordEndpoint returns an endpoint that invokes ChangePassword on the service.
func MakeChangePasswordEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangePasswordRequest)
		success, err := s.ChangePassword(ctx, req.Login, req.Curr, req.New)
		return &ChangePasswordResponse{
			Err:     err,
			Success: success,
		}, nil
	}
}

// Failed implements Failer.
func (r ChangePasswordResponse) Failed() error {
	return r.Err
}

// SignUpRequest collects the request parameters for the SignUp method.
type SignUpRequest struct {
	User model.User `json:"user"`
}

// SignUpResponse collects the response parameters for the SignUp method.
type SignUpResponse struct {
	Id  string `json:"id"`
	Err error  `json:"err"`
}

// MakeSignUpEndpoint returns an endpoint that invokes SignUp on the service.
func MakeSignUpEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		id, err := s.SignUp(ctx, req.User)
		return &SignUpResponse{
			Err: err,
			Id:  id,
		}, nil
	}
}

// Failed implements Failer.
func (r SignUpResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Err error `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.Id)
		return &DeleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Auth implements Service. Primarily useful in a client.
func (e Endpoints) Auth(ctx context.Context, login string, password string) (jwt string, err error) {
	request := AuthRequest{
		Login:    login,
		Password: password,
	}
	response, err := e.AuthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(*AuthResponse).Jwt, response.(*AuthResponse).Err
}

// ChangePassword implements Service. Primarily useful in a client.
func (e Endpoints) ChangePassword(ctx context.Context, login string, curr string, new string) (success bool, err error) {
	request := ChangePasswordRequest{
		Curr:  curr,
		Login: login,
		New:   new,
	}
	response, err := e.ChangePasswordEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(*ChangePasswordResponse).Success, response.(*ChangePasswordResponse).Err
}

// SignUp implements Service. Primarily useful in a client.
func (e Endpoints) SignUp(ctx context.Context, user model.User) (id string, err error) {
	request := SignUpRequest{User: user}
	response, err := e.SignUpEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(*SignUpResponse).Id, response.(*SignUpResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (err error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(*DeleteResponse).Err
}
