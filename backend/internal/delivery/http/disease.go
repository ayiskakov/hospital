package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/khanfromasia/hospital/backend/internal/storage"
)

type diseaseTypeCreateRequest struct {
	Description string `json:"description"`
}

type diseaseTypeCreateResponse struct {
	DiseaseType entity.DiseaseType `json:"disease_type"`
}

// diseaseTypeCreate performs the disease type create operation http.Handler.
func (h *handler) diseaseTypeCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req diseaseTypeCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var diseaseType entity.DiseaseType

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			diseaseType, err = q.DiseaseTypeCreate(r.Context(), entity.DiseaseType{
				Description: req.Description,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseTypeCreate():", err)
			return
		}

		respondJSON(w, http.StatusOK, diseaseTypeCreateResponse{DiseaseType: diseaseType})
	})
}

type diseaseTypeGetAllResponse struct {
	DiseaseTypes []entity.DiseaseType `json:"disease_types"`
}

// diseaseTypeGetAll performs the disease type get all operation http.Handler.
func (h *handler) diseaseTypeGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var diseaseTypes []entity.DiseaseType

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			diseaseTypes, err = q.DiseaseTypeGetAll(r.Context())

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseTypeGetAll():", err)
			return
		}
		respondJSON(w, http.StatusOK, diseaseTypeGetAllResponse{DiseaseTypes: diseaseTypes})
	})
}

type diseaseTypeUpdateRequest struct {
	Description *string `json:"description"`
}

type diseaseTypeUpdateResponse struct {
	DiseaseType entity.DiseaseType `json:"disease_type"`
}

// diseaseTypeUpdate performs the disease type update operation http.Handler.
func (h *handler) diseaseTypeUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, errP := strconv.ParseInt(httprouter.ParamsFromContext(r.Context()).ByName("id"), 10, 64)

		if errP != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var req diseaseTypeUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var diseaseType entity.DiseaseType

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			diseaseType, err = q.DiseaseTypeUpdate(r.Context(), id, entity.DiseaseTypeUpdate{
				Description: req.Description,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseTypeUpdate():", err)
			return
		}

		respondJSON(w, http.StatusOK, diseaseTypeUpdateResponse{DiseaseType: diseaseType})
	})
}

// diseaseTypeDelete performs the disease type delete operation http.Handler.
func (h *handler) diseaseTypeDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, errP := strconv.ParseInt(httprouter.ParamsFromContext(r.Context()).ByName("id"), 10, 64)

		if errP != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			return q.DiseaseTypeDelete(r.Context(), id)
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseTypeDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}

type diseaseCreateRequest struct {
	ID          int64  `json:"id"`
	DiseaseCode string `json:"disease_code"`
	Description string `json:"description"`
	Pathogen    string `json:"pathogen"`
}

type diseaseCreateResponse struct {
	Disease entity.Disease `json:"disease"`
}

// diseaseCreate performs the disease create operation http.Handler.
func (h *handler) diseaseCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req diseaseCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var disease entity.Disease

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			disease, err = q.DiseaseCreate(r.Context(), entity.Disease{
				DiseaseType: entity.DiseaseType{ID: req.ID},
				DiseaseCode: req.DiseaseCode,
				Description: req.Description,
				Pathogen:    req.Pathogen,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseCreate():", err)
			return
		}

		respondJSON(w, http.StatusOK, diseaseCreateResponse{Disease: disease})
	})
}

type diseaseGetAllResponse struct {
	Diseases []entity.Disease `json:"diseases"`
}

// diseaseGetAll performs the disease get all operation http.Handler.
func (h *handler) diseaseGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var diseases []entity.Disease

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			diseases, err = q.DiseaseGetAll(r.Context())

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseGetAll():", err)
			return
		}
		respondJSON(w, http.StatusOK, diseaseGetAllResponse{Diseases: diseases})
	})
}

type diseaseUpdateRequest struct {
	ID          *int64  `json:"id"`
	DiseaseCode *string `json:"disease_code"`
	Description *string `json:"description"`
	Pathogen    *string `json:"pathogen"`
}

type diseaseUpdateResponse struct {
	Disease entity.Disease `json:"disease"`
}

// diseaseUpdate performs the disease update operation http.Handler.
func (h *handler) diseaseUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		diseaseCode := httprouter.ParamsFromContext(r.Context()).ByName("disease_code")

		var req diseaseUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var disease entity.Disease

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			disease, err = q.DiseaseUpdate(r.Context(), diseaseCode, entity.DiseaseUpdate{
				ID:          req.ID,
				DiseaseCode: req.DiseaseCode,
				Description: req.Description,
				Pathogen:    req.Pathogen,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseUpdate():", err)
			return
		}

		respondJSON(w, http.StatusOK, diseaseUpdateResponse{Disease: disease})
	})
}

// diseaseDelete performs the disease delete operation http.Handler.
func (h *handler) diseaseDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		diseaseCode := httprouter.ParamsFromContext(r.Context()).ByName("disease_code")

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			return q.DiseaseDelete(r.Context(), diseaseCode)
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, "internal server error")
			log.Println("diseaseDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}
