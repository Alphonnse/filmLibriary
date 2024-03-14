package model

type Register struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
