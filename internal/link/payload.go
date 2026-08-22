package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LoginResponce struct {
	Token string `json:"token"`
}


type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponce struct {
	Token string `json:"token"`
}
