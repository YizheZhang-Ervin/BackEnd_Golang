package structure

type DeleteProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteProfileResponse) error() error { return r.Err }

type GetProfileResponse struct {
	Profile Profile `json:"profile,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func (r GetProfileResponse) error() error { return r.Err }

type PatchProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r PatchProfileResponse) error() error { return r.Err }

type PostProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r PostProfileResponse) error() error { return r.Err }

type PutProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r PutProfileResponse) error() error { return nil }
