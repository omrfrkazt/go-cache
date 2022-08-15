package main

type BaseResponseModel struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type UserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SecretModel struct {
	UserName string `json:"username"`
	Key string `json:"key"`
}

