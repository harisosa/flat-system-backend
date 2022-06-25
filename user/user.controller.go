package user

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/harisosa/flat-system-backend/entity"
	"github.com/harisosa/flat-system-backend/helper"
)

//userSettingController  represent the httphandler for user
type userController struct {
	usc    entity.UserUsecase
	logger *log.Logger
	help   helper.Helper
}

// NewUserController will initialize the user endpoint
func NewUserController(router *mux.Router, usc entity.UserUsecase, help helper.Helper) {
	userHandler := &userController{
		logger: log.New(os.Stdout, "User Controller ", log.LstdFlags),
		usc:    usc,
		help:   help,
	}
	userRouter := router.PathPrefix("/api/user").Subrouter()
	userRouter.HandleFunc("", userHandler.getall).Methods("GET")
	userRouter.HandleFunc("/submit", userHandler.submit).Methods("POST")
	userRouter.HandleFunc("/remove/{id}", userHandler.remove).Methods("DELETE")

}

func (ctrl *userController) getall(w http.ResponseWriter, r *http.Request) {
	res, err := ctrl.usc.GetAll()
	if err != nil {
		ctrl.logger.Printf("Error on getByID: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.help.ResponseHandler(w, res, "", http.StatusOK)
}
func (ctrl *userController) submit(w http.ResponseWriter, r *http.Request) {
	var usr entity.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		ctrl.logger.Printf("Error on submit: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ctrl.usc.Upsert(usr)
	if err != nil {
		ctrl.logger.Printf("Error on submit: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.help.ResponseHandler(w, nil, "Submit Sucessful ", http.StatusOK)
}
func (ctrl *userController) remove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]
	err := ctrl.usc.Remove(userId)
	if err != nil {
		ctrl.logger.Printf("Error on submit: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.help.ResponseHandler(w, nil, "Delete User Sucessful ", http.StatusOK)
}
