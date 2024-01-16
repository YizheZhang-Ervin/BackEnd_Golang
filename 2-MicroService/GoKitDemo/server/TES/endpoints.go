package TES

import (
	"context"
	"net/url"
	"server/TES/structure"
	"server/TES/utils"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Endpoints struct {
	PostProfileEndpoint   endpoint.Endpoint
	GetProfileEndpoint    endpoint.Endpoint
	PutProfileEndpoint    endpoint.Endpoint
	PatchProfileEndpoint  endpoint.Endpoint
	DeleteProfileEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostProfileEndpoint:   MakePostProfileEndpoint(s),
		GetProfileEndpoint:    MakeGetProfileEndpoint(s),
		PutProfileEndpoint:    MakePutProfileEndpoint(s),
		PatchProfileEndpoint:  MakePatchProfileEndpoint(s),
		DeleteProfileEndpoint: MakeDeleteProfileEndpoint(s),
	}
}

func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		PostProfileEndpoint:   httptransport.NewClient("POST", tgt, utils.EncodePostProfileRequest, utils.DecodePostProfileResponse, options...).Endpoint(),
		GetProfileEndpoint:    httptransport.NewClient("GET", tgt, utils.EncodeGetProfileRequest, utils.DecodeGetProfileResponse, options...).Endpoint(),
		PutProfileEndpoint:    httptransport.NewClient("PUT", tgt, utils.EncodePutProfileRequest, utils.DecodePutProfileResponse, options...).Endpoint(),
		PatchProfileEndpoint:  httptransport.NewClient("PATCH", tgt, utils.EncodePatchProfileRequest, utils.DecodePatchProfileResponse, options...).Endpoint(),
		DeleteProfileEndpoint: httptransport.NewClient("DELETE", tgt, utils.EncodeDeleteProfileRequest, utils.DecodeDeleteProfileResponse, options...).Endpoint(),
	}, nil
}

// PostProfile implements Service. Primarily useful in a client.
func (e Endpoints) PostProfile(ctx context.Context, p structure.Profile) error {
	request := structure.PostProfileRequest{Profile: p}
	response, err := e.PostProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(structure.PostProfileResponse)
	return resp.Err
}

// GetProfile implements Service. Primarily useful in a client.
func (e Endpoints) GetProfile(ctx context.Context, id string) (structure.Profile, error) {
	request := structure.GetProfileRequest{ID: id}
	response, err := e.GetProfileEndpoint(ctx, request)
	if err != nil {
		return structure.Profile{}, err
	}
	resp := response.(structure.GetProfileResponse)
	return resp.Profile, resp.Err
}

// PutProfile implements Service. Primarily useful in a client.
func (e Endpoints) PutProfile(ctx context.Context, id string, p structure.Profile) error {
	request := structure.PutProfileRequest{ID: id, Profile: p}
	response, err := e.PutProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(structure.PutProfileResponse)
	return resp.Err
}

// PatchProfile implements Service. Primarily useful in a client.
func (e Endpoints) PatchProfile(ctx context.Context, id string, p structure.Profile) error {
	request := structure.PatchProfileRequest{ID: id, Profile: p}
	response, err := e.PatchProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(structure.PatchProfileResponse)
	return resp.Err
}

// DeleteProfile implements Service. Primarily useful in a client.
func (e Endpoints) DeleteProfile(ctx context.Context, id string) error {
	request := structure.DeleteProfileRequest{ID: id}
	response, err := e.DeleteProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(structure.DeleteProfileResponse)
	return resp.Err
}

// MakePostProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePostProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(structure.PostProfileRequest)
		e := s.PostProfile(ctx, req.Profile)
		return structure.PostProfileResponse{Err: e}, nil
	}
}

// MakeGetProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(structure.GetProfileRequest)
		p, e := s.GetProfile(ctx, req.ID)
		return structure.GetProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakePutProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePutProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(structure.PutProfileRequest)
		e := s.PutProfile(ctx, req.ID, req.Profile)
		return structure.PutProfileResponse{Err: e}, nil
	}
}

// MakePatchProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePatchProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(structure.PatchProfileRequest)
		e := s.PatchProfile(ctx, req.ID, req.Profile)
		return structure.PatchProfileResponse{Err: e}, nil
	}
}

// MakeDeleteProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeDeleteProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(structure.DeleteProfileRequest)
		e := s.DeleteProfile(ctx, req.ID)
		return structure.DeleteProfileResponse{Err: e}, nil
	}
}
