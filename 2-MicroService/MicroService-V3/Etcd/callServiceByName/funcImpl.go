package serivces

import (
	"context"
	"net/http"
	"strconv"
)

func ProdEncodeFunc(ctx context.Context, httpRequest *http.Request, requestParams interface{}) error {
	prodr := requestParams.(ProdRequest)
	httpRequest.URL.Path += "/product/" + strconv.Itoa(prodr.ProdId)
	return nil
}
