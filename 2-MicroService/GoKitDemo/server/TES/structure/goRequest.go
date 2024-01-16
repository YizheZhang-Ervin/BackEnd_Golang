package structure

type DeleteProfileRequest struct {
	ID string
}

type GetProfileRequest struct {
	ID string
}

type PatchProfileRequest struct {
	ID      string
	Profile Profile
}

type PostProfileRequest struct {
	Profile Profile
}

type PutProfileRequest struct {
	ID      string
	Profile Profile
}
