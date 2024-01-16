package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/TES/structure"
)

// 通用请求编码=================================
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

// 请求编码=================================
func EncodePostProfileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("POST").Path("/profiles/")
	req.URL.Path = "/profiles/"
	return encodeRequest(ctx, req, request)
}

func EncodeGetProfileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/profiles/{id}")
	r := request.(structure.GetProfileRequest)
	profileID := url.QueryEscape(r.ID)
	req.URL.Path = "/profiles/" + profileID
	return encodeRequest(ctx, req, request)
}

func EncodePutProfileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("PUT").Path("/profiles/{id}")
	r := request.(structure.PutProfileRequest)
	profileID := url.QueryEscape(r.ID)
	req.URL.Path = "/profiles/" + profileID
	return encodeRequest(ctx, req, request)
}

func EncodePatchProfileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("PATCH").Path("/profiles/{id}")
	r := request.(structure.PatchProfileRequest)
	profileID := url.QueryEscape(r.ID)
	req.URL.Path = "/profiles/" + profileID
	return encodeRequest(ctx, req, request)
}

func EncodeDeleteProfileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("DELETE").Path("/profiles/{id}")
	r := request.(structure.DeleteProfileRequest)
	profileID := url.QueryEscape(r.ID)
	req.URL.Path = "/profiles/" + profileID
	return encodeRequest(ctx, req, request)
}
