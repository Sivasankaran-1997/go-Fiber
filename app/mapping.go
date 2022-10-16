package app

import "fiberscurd/controller"

func Routers() {
	r.Post("/users/create", controller.CreateUser)
	r.Get("/users/login", controller.Login)
	r.Get("/users/getbyemail", controller.GetUserByEmail)
	r.Delete("/users/deletebyemail", controller.DeleteUserByEmail)
	r.Patch("/users/updatebyemail", controller.UpdateUserByEmail)
}
