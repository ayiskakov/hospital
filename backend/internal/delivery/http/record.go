package http

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/khanfromasia/hospital/backend/internal/storage"
)

type recordCreateRequest struct {
	Email         string `json:"email"`
	Cname         string `json:"cname"`
	DiseaseCode   string `json:"disease_code"`
	TotalDeaths   int64  `json:"total_deaths"`
	TotalPatients int64  `json:"total_patients"`
}

type recordCreateResponse struct {
	Record entity.Record `json:"record"`
}

// recordCreate performs the record create operation http.Handler.
func (h *handler) recordCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req recordCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}
		var record entity.Record

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			record, err = q.RecordCreate(r.Context(), entity.Record{
				PublicServant: entity.PublicServant{
					User: entity.User{
						Email: req.Email,
					},
				},
				Country: entity.Country{
					Cname: req.Cname,
				},
				Disease: entity.Disease{
					DiseaseCode: req.DiseaseCode,
				},
				TotalDeaths:   req.TotalDeaths,
				TotalPatients: req.TotalPatients,
			})

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("recordCreate():", err)
			return
		}

		respondJSON(w, http.StatusOK, recordCreateResponse{
			Record: record,
		})
	})
}

type recordGetAllResponse struct {
	Records []entity.Record `json:"records"`
}

// recordGetAll performs the record get all operation http.Handler.
func (h *handler) recordGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var records []entity.Record

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			records, err = q.RecordGetAll(r.Context())

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("recordGetAll():", err)
			return
		}

		respondJSON(w, http.StatusOK, recordGetAllResponse{
			Records: records,
		})
	})
}

func (h *handler) recordDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cname := httprouter.ParamsFromContext(r.Context()).ByName("cname")
		diseaseCode := httprouter.ParamsFromContext(r.Context()).ByName("disease_code")
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			err = q.RecordDelete(r.Context(), cname, email, diseaseCode)

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("recordDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}

type recordUpdateRequest struct {
	TotalDeaths   *int64 `json:"total_deaths"`
	TotalPatients *int64 `json:"total_patients"`
}

type recordUpdateResponse struct {
	Record entity.Record `json:"record"`
}

func (h *handler) recordUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req recordUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		cname := httprouter.ParamsFromContext(r.Context()).ByName("cname")
		diseaseCode := httprouter.ParamsFromContext(r.Context()).ByName("disease_code")
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		var record entity.Record

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			record, err = q.RecordUpdate(r.Context(), cname, email, diseaseCode, entity.RecordUpdate{
				TotalDeaths:   req.TotalDeaths,
				TotalPatients: req.TotalPatients,
			})

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("recordUpdate():", err)
			return
		}

		respondJSON(w, http.StatusOK, recordUpdateResponse{
			Record: record,
		})
	})
}
