package http

import (
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/khanfromasia/hospital/backend/internal/storage"
)

type discoveryCreateRequest struct {
	Cname        string    `json:"cname"`
	DiseaseCode  string    `json:"disease_code"`
	FirstEncDate time.Time `json:"first_enc_date"`
}

type discoveryCreateResponse struct {
	Discovery entity.Discovery `json:"discovery"`
}

func (h *handler) discoveryCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req discoveryCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var discovery entity.Discovery

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			discovery, err = q.DiscoveryCreate(r.Context(), entity.Discovery{
				Country: entity.Country{
					Cname: req.Cname,
				},
				Disease: entity.Disease{
					DiseaseCode: req.DiseaseCode,
				},
				FirstEncDate: req.FirstEncDate,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("discoveryCreate():", err)
			return
		}

		respondJSON(w, http.StatusOK, discoveryCreateResponse{Discovery: discovery})
	})
}

type discoveryGetAllResponse struct {
	Discoveries []entity.Discovery `json:"discoveries"`
}

func (h *handler) discoveryGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var discoveries []entity.Discovery

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			discoveries, err = q.DiscoveryGetAll(r.Context())

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("discoveryGetAll():", err)
			return
		}

		respondJSON(w, http.StatusOK, discoveryGetAllResponse{Discoveries: discoveries})
	})
}

func (h *handler) discoveryDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cname := httprouter.ParamsFromContext(r.Context()).ByName("cname")
		diseaseCode := httprouter.ParamsFromContext(r.Context()).ByName("disease_code")

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			err = q.DiscoveryDelete(r.Context(), cname, diseaseCode)

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("discoveryDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}

type discoveryUpdateRequest struct {
	Cname        *string    `json:"cname"`
	DiseaseCode  *string    `json:"disease_code"`
	FirstEncDate *time.Time `json:"first_enc_date"`
}

type discoveryUpdateResponse struct {
	Discovery entity.Discovery `json:"discovery"`
}

func (h *handler) discoveryUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cname := httprouter.ParamsFromContext(r.Context()).ByName("cname")
		diseaseCode := httprouter.ParamsFromContext(r.Context()).ByName("disease_code")

		var req discoveryUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var discovery entity.Discovery

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			discovery, err = q.DiscoveryUpdate(r.Context(), cname, diseaseCode, entity.DiscoveryUpdate{
				Cname:        req.Cname,
				DiseaseCode:  req.DiseaseCode,
				FirstEncDate: req.FirstEncDate,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("discoveryUpdate():", err)
			return
		}

		respondJSON(w, http.StatusOK, discoveryUpdateResponse{Discovery: discovery})
	})
}
