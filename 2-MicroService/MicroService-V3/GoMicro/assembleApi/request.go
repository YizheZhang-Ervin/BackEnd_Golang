package sidecar

type JSONRequest struct {
	Jsonrpc string
	Method  string
	Params  []*Service
	Id      int
}

func NewJSONRequest(service *Service, endpoint string) *JSONRequest {
	return &JSONRequest{Jsonrpc: "2.0", Method: endpoint, Params: []*Service{service}, Id: 1}
}
