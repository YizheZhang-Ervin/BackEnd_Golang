package utils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"server/TES/structure"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// 请求解码=================================
func DecodePostProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req structure.PostProfileRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Profile); e != nil {
		return nil, e
	}
	return req, nil
}

func DecodeGetProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return structure.GetProfileRequest{ID: id}, nil
}

func DecodePutProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	var profile structure.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return structure.PutProfileRequest{
		ID:      id,
		Profile: profile,
	}, nil
}

func DecodePatchProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	var profile structure.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return structure.PatchProfileRequest{
		ID:      id,
		Profile: profile,
	}, nil
}

func DecodeDeleteProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return structure.DeleteProfileRequest{ID: id}, nil
}

// 响应解码=================================
func DecodePostProfileResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response structure.PostProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func DecodeGetProfileResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response structure.GetProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func DecodePutProfileResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response structure.PutProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func DecodePatchProfileResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response structure.PatchProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func DecodeDeleteProfileResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response structure.DeleteProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
