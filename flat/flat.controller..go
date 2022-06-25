package flat

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
type flatController struct {
	fsc    entity.FlateUsecase
	logger *log.Logger
	help   helper.Helper
}

// NewUserController will initialize the user endpoint
func NewFlatController(router *mux.Router, fsc entity.FlateUsecase, help helper.Helper) {
	flatHandler := &flatController{
		logger: log.New(os.Stdout, "User Setting Controller ", log.LstdFlags),
		fsc:    fsc,
		help:   help,
	}
	flatRouter := router.PathPrefix("/api/flat").Subrouter()
	flatRouter.HandleFunc("", flatHandler.getall).Methods("GET")
	flatRouter.HandleFunc("/submit", flatHandler.submit).Methods("POST")
}

func (ctrl *flatController) getall(w http.ResponseWriter, r *http.Request) {
	res, err := ctrl.fsc.GetAll()
	if err != nil {
		ctrl.logger.Printf("Error on getByID: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.help.ResponseHandler(w, res, "", http.StatusOK)
}
func (ctrl *flatController) submit(w http.ResponseWriter, r *http.Request) {
	var fs entity.Flat
	err := json.NewDecoder(r.Body).Decode(&fs)
	if err != nil {
		ctrl.logger.Printf("Error on submit: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ctrl.fsc.Upsert(fs)
	if err != nil {
		ctrl.logger.Printf("Error on submit: %s\n", err)
		ctrl.help.ResponseHandler(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}
	ctrl.help.ResponseHandler(w, nil, "Submit Sucessful ", http.StatusOK)
}
