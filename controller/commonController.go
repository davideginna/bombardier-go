package controller

import "log"

type CommonController struct {
}

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json; charset=UTF-8"
)

func CommonControllerFactory() CommonController {
	log.Println("CommonControllerFactory -- START --")
	var commonController = CommonController{}
	return commonController
}
