package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/internal/storage"
)

const (
	v0 = "/api/v0/"
)

type handler struct {
	storage *storage.Storage // sorry for this
}

func NewHandler(storage *storage.Storage) *handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) SetupRoutes() *httprouter.Router {
	router := httprouter.New()

	router.Handler(http.MethodGet, v0+"countries", h.countryGetAll())
	router.Handler(http.MethodPost, v0+"countries", h.countryCreate())
	router.Handler(http.MethodPut, v0+"countries/:cname", h.countryUpdate())
	router.Handler(http.MethodDelete, v0+"countries/:cname", h.countryDelete())

	router.Handler(http.MethodGet, v0+"users", h.userGetAll())
	router.Handler(http.MethodPost, v0+"users", h.userCreate())
	router.Handler(http.MethodPut, v0+"users/:email", h.userUpdate())
	router.Handler(http.MethodDelete, v0+"users/:email", h.userDelete())

	router.Handler(http.MethodGet, v0+"doctors", h.doctorGetAll())
	router.Handler(http.MethodPost, v0+"doctors", h.doctorCreate())
	router.Handler(http.MethodPut, v0+"doctors/:email", h.doctorUpdate())
	router.Handler(http.MethodDelete, v0+"doctors/:email", h.doctorDelete())

	router.Handler(http.MethodGet, v0+"publicServants", h.publicServantGetAll())
	router.Handler(http.MethodPost, v0+"publicServants", h.publicServantCreate())
	router.Handler(http.MethodPut, v0+"publicServants/:email", h.publicServantUpdate())
	router.Handler(http.MethodDelete, v0+"publicServants/:email", h.publicServantDelete())

	router.Handler(http.MethodGet, v0+"diseaseTypes", h.diseaseTypeGetAll())
	router.Handler(http.MethodPost, v0+"diseaseTypes", h.diseaseTypeCreate())
	router.Handler(http.MethodPut, v0+"diseaseTypes/:id", h.diseaseTypeUpdate())
	router.Handler(http.MethodDelete, v0+"diseaseTypes/:id", h.diseaseTypeDelete())

	router.Handler(http.MethodGet, v0+"diseases", h.diseaseGetAll())
	router.Handler(http.MethodPost, v0+"diseases", h.diseaseCreate())
	router.Handler(http.MethodPut, v0+"diseases/:id", h.diseaseUpdate())
	router.Handler(http.MethodDelete, v0+"diseases/:id", h.diseaseDelete())

	router.Handler(http.MethodGet, v0+"discovery", h.discoveryGetAll())
	router.Handler(http.MethodPost, v0+"discovery", h.discoveryCreate())
	router.Handler(http.MethodPut, v0+"discovery/:cname/:disease_code", h.discoveryUpdate())
	router.Handler(http.MethodDelete, v0+"discovery/:cname/:disease_code", h.discoveryDelete())

	router.Handler(http.MethodGet, v0+"records", h.recordGetAll())
	router.Handler(http.MethodPost, v0+"records", h.recordCreate())
	router.Handler(http.MethodPut, v0+"records/:id", h.recordUpdate())
	router.Handler(http.MethodDelete, v0+"records/:id", h.recordDelete())

	router.Handler(http.MethodGet, v0+"specialize", h.specializeGetAll())
	router.Handler(http.MethodPost, v0+"specialize", h.specializeCreate())
	router.Handler(http.MethodPut, v0+"specialize/:email/:id", h.specializeDelete())

	return router
}
