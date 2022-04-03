package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.tdnet.it/cochise/golang/template-project-ms/pkg/template"
	"gitlab.tdnet.it/cochise/golang/template-project-ms/pkg/utils"
)

type TemplateAPI struct {
	templateService template.Service
}

func TemplateAPIFactory() TemplateAPI {
	var api = TemplateAPI{template.ServiceFactory()}
	return api
}

func (c TemplateAPI) RecuperaTemplateByID(w http.ResponseWriter, r *http.Request) {
	log.Println("TemplateApi:: servizio RecuperaTemplateById")
	user := r.Header.Get("X_RTSYSID")
	log.Println("RecuperaTemplateById:: user=" + user)
	pathParams := mux.Vars(r)
	log.Println("id=" + pathParams["id"])
	dto, err := c.templateService.RecuperaTemplateByID(pathParams["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, dto)
}
