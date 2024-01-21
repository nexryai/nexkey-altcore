package schema

type SignInRequest struct {
	Password  string `json:"password" xml:"password" form:"password"`
	Username  string `json:"username" xml:"username" form:"username"`
	Hcaptcha  string `json:"hcaptcha-response"`
	Recaptcha string `json:"g-recaptcha-response"`
}

type SignInResp struct {
	Token  string `json:"i"`
	UserId string `json:"userId"`
}
