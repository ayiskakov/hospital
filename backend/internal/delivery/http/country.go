package http

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/khanfromasia/hospital/backend/internal/storage"
)

type countryCreateRequest struct {
	Cname      string `json:"cname"`
	Population int64  `json:"population"`
}

type countryCreateResponse struct {
	Country entity.Country `json:"country"`
}

func (h *handler) countryCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req countryCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var country entity.Country

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			country, err = q.CountryCreate(r.Context(), entity.Country{
				Cname:      req.Cname,
				Population: req.Population,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("countryCreate():", err)
			return
		}

		respondJSON(w, http.StatusOK, countryCreateResponse{Country: country})
	})
}

type countryGetAllResponse struct {
	Countries []entity.Country `json:"countries"`
}

func (h *handler) countryGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var countries []entity.Country

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			countries, err = q.CountryGetAll(r.Context())

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("countryGetAll():", err)
			return
		}

		respondJSON(w, http.StatusOK, countryGetAllResponse{Countries: countries})
	})
}

func (h *handler) countryDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cname := httprouter.ParamsFromContext(r.Context()).ByName("cname")

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			return q.CountryDelete(r.Context(), cname)
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("countryDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}

type countryUpdateRequest struct {
	Population int64 `json:"population"`
}

type countryUpdateResponse struct {
	Country entity.Country `json:"country"`
}

func (h *handler) countryUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cname := httprouter.ParamsFromContext(r.Context()).ByName("cname")

		var req countryUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var country entity.Country

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			country, err = q.CountryUpdate(r.Context(), entity.Country{
				Cname:      cname,
				Population: req.Population,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("countryUpdate():", err)
			return
		}

		respondJSON(w, http.StatusOK, countryUpdateResponse{Country: country})
	})
}
