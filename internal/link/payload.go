package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkGoToRequest struct {
	Hash string `json:"hash" validate:"required,string"`
}