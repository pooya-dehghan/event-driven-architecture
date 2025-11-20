package userhandler

import userservice "github.com/pooya/services"

type Handler struct {
	userSvc userservice.Service
}

func New(userSvc userservice.Service) Handler {

	return Handler{userSvc: userSvc}
}
