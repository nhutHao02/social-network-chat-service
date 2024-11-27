package model

type UserInfo struct {
	ID       int     `json:"id"`
	Email    *string `json:"email"`
	FullName *string `json:"fullName"`
	UrlAvt   *string `json:"urlAvt"`
}
