package template

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate genny -in=template-service.go -out=service.go gen "Template=${TIPO}"

// Service interfaccia servizio
type Service interface {
	CreaTemplate(templateDTO TemplateDTO) (string, error)
	RecuperaTemplateByID(id string) (TemplateDTO, error)
	ListaTemplate() ([]TemplateDTO, error)
	AggiornaTemplateByID(id string, templateDTO TemplateDTO) error
	CancellaTemplateByID(id string) error
	RecuperaByFiltro(filtro FiltroDTO) ([]TemplateDTO, error)
}

// ServiceFactory factory
func ServiceFactory() Service {
	dao := daoFactory()
	return templateServiceImpl{templateDao: dao}
}

// templateServiceImpl ..
type templateServiceImpl struct {
	templateDao dao
}

// AggiornaTemplateByID ..
func (ms templateServiceImpl) AggiornaTemplateByID(id string, templateDTO TemplateDTO) error {
	log.Println("TemplateServiceImpl:: AggiornaTemplateByID: ", id)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	templateDB, _ := ms.templateDao.RecuperaTemplateByID(id)
	var dataCreazione = templateDB.DataCreazione
	template := dataCreazioneTemplateDtoToEntity(templateDTO, dataCreazione)
	template.ID = objectID
	if err != nil {
		return err
	}
	err = ms.templateDao.AggiornaTemplateByID(template)
	return err
}

// RecuperaTemplateByID ..
func (ms templateServiceImpl) RecuperaTemplateByID(id string) (TemplateDTO, error) {
	log.Println("TemplateServiceImpl:: RecuperaTemplateByID: ", id)
	if id != "" {
		template, err := ms.templateDao.RecuperaTemplateByID(id)
		if err != nil {
			return TemplateDTO{}, err
		}
		templateDTO := templateToDTO(template)
		return templateDTO, err
	}
	return TemplateDTO{}, errors.New("id non valorizzato")
}

// ListaTemplate ..
func (ms templateServiceImpl) ListaTemplate() ([]TemplateDTO, error) {
	log.Println("TemplateServiceImpl:: ListaTemplate")
	lista, err := ms.templateDao.ListaTemplate()
	listaDTO := templateListToDTO(lista)
	return listaDTO, err
}

// CreaTemplate ..
func (ms templateServiceImpl) CreaTemplate(templateDTO TemplateDTO) (string, error) {
	log.Println("TemplateServiceImpl:: CreaTemplate")
	inserito, err := ms.templateDao.CreaTemplate(templateFromDTO(templateDTO))
	return inserito, err
}

// CancellaTemplateByID ..
func (ms templateServiceImpl) CancellaTemplateByID(id string) error {
	log.Println("TemplateServiceImpl:: CancellaTemplateByID: ", id)
	if id != "" {
		err := ms.templateDao.CancellaTemplateByID(id)
		return err
	}
	return errors.New("id Template non valorizzato")
}

// RecuperaByFiltro ..
func (ms templateServiceImpl) RecuperaByFiltro(filtro FiltroDTO) ([]TemplateDTO, error) {
	log.Println("TemplateServiceImpl:: RecuperaByFiltro: ", filtro)
	if filtro.Attributo != "" || filtro.Codice != 0 {
		listaTemplate, err := ms.templateDao.ListaFiltroTemplate(filtro)
		log.Println("RecuperoListaTemplate: ", listaTemplate)
		if err != nil {
			return []TemplateDTO{}, err
		}
		//lista, err := ms.TemplateDao.Lista()
		listaDTO := templateListToDTO(listaTemplate)
		log.Println("ListaTemplate:", listaTemplate)
		return listaDTO, err
	}
	return []TemplateDTO{}, errors.New("Non valorizzato")
}
