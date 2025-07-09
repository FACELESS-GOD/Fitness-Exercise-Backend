package Router

import (
	"github.com/FACELESS-GOD/Fitness-Exercise-Backend.git/Packages/Controller"
	"github.com/gorilla/mux"
)

type CustomRouter struct {
	Router *mux.Router
	Ctrl   *Controller.ControllerStruct
}

func NewCustomRouter(Router *mux.Router, Ctrl *Controller.ControllerStruct) *CustomRouter {
	currRouter := CustomRouter{}
	currRouter.Ctrl = Ctrl
	currRouter.Router = Router
	return &currRouter
}

func (CustomRouter *CustomRouter) UpdateCustomRouter() {
	CustomRouter.Router.HandleFunc("/add", CustomRouter.Ctrl.AddUser).Methods("GET")
	CustomRouter.Router.HandleFunc("/loging", CustomRouter.Ctrl.ValidateUser).Methods("GET")
}
