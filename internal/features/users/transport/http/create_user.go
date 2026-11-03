package users_transport_http

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Name        string  `json:"name"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

type CreateUserResponse struct {
	ID          int64  `json:"id"`
	Version     int    `json:"version"`
	Name        string `json:"name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if  err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


}
