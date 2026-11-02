package users_transport_http

type UsersHTTPHandler struct {
	usersService UserService
}

type UserService interface {

}

func NewUsersHTTPHandler(usersService UserService) *UsersHTTPHandler {
	return &UsersHTTPHandler{usersService: usersService}
}
