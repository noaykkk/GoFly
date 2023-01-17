// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
