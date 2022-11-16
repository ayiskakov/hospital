package http

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/internal/entity"
	"github.com/khanfromasia/hospital/internal/storage"
)

type specializeCreateRequest struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type specializeCreateResponse struct {
	Specialize entity.Specialize `json:"specialize"`
}

// specializeCreate performs the specialize create operation http.Handler.
func (h *handler) specializeCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req specializeCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}
		var specialize entity.Specialize

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			specialize, err = q.SpecializeCreate(r.Context(), entity.Specialize{
				DiseaseType: entity.DiseaseType{ID: req.ID},
				Doctor:      entity.Doctor{User: entity.User{Email: req.Email}},
			})

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("specializeCreate():", err)
			return
		}

		respondJSON(w, http.StatusOK, specializeCreateResponse{
			Specialize: specialize,
		})
	})
}

// specializeDelete performs the specialize delete operation http.Handler.
func (h *handler) specializeDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")
		id := httprouter.ParamsFromContext(r.Context()).ByName("id")

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			err = q.SpecializeDelete(r.Context(), email, id)

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("specializeDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}

type specializeGetAllResponse struct {
	Specializes []entity.Specialize `json:"specializes"`
}

// specializeGetAll performs the specialize get all operation http.Handler.
func (h *handler) specializeGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var specializes []entity.Specialize

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			specializes, err = q.SpecializeGetAll(r.Context())

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("specializeGetAll():", err)
			return
		}

		respondJSON(w, http.StatusOK, specializeGetAllResponse{
			Specializes: specializes,
		})
	})
}
