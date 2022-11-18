package http

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/khanfromasia/hospital/backend/internal/storage"
)

type publicServantCreateRequest struct {
	User struct {
		//Name    string `json:"name"`
		//Surname string `json:"surname"`
		Email string `json:"email"`
		//Phone   string `json:"phone"`
		//Salary  int64  `json:"salary"`
		//Country struct {
		//	Cname string `json:"cname"`
		//} `json:"country"`
	} `json:"user"`
	PublicServant struct {
		Department string `json:"department"`
	} `json:"public_servant"`
}

type publicServantCreateResponse struct {
	PublicServant entity.PublicServant `json:"public_servant"`
}

func (h *handler) publicServantCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req publicServantCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var publicServant entity.PublicServant

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			//publicServant.User, err = q.UserCreate(r.Context(), entity.User{
			//	//Name:    req.User.Name,
			//	//Surname: req.User.Surname,
			//	Email:   req.User.Email,
			//	//Phone:   req.User.Phone,
			//	//Salary:  req.User.Salary,
			//	Country: entity.Country{
			//		Cname: req.User.Country.Cname,
			//	},
			//})
			publicServant.User.Email = req.User.Email
			publicServant.Department = req.PublicServant.Department

			publicServant, err = q.PublicServantCreate(r.Context(), publicServant)

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("publicServantCreate():", err)
			return
		}

		respondJSON(w, http.StatusCreated, publicServantCreateResponse{
			PublicServant: publicServant,
		})
	})
}

type publicServantGetAllResponse struct {
	PublicServants []entity.PublicServant `json:"public_servants"`
}

func (h *handler) publicServantGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var publicServants []entity.PublicServant

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			publicServants, err = q.PublicServantGetAll(r.Context())

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("publicServantGetAll():", err)
			return
		}

		respondJSON(w, http.StatusOK, publicServantGetAllResponse{
			PublicServants: publicServants,
		})
	})
}

type publicServantUpdateRequest struct {
	Department *string `json:"department"`
}

type publicServantUpdateResponse struct {
	PublicServant entity.PublicServant `json:"public_servant"`
}

func (h *handler) publicServantUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		var req publicServantUpdateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var publicServant entity.PublicServant

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			if req.Department != nil {
				publicServant.Department = *req.Department
			}

			publicServant, err = q.PublicServantUpdate(r.Context(), email, entity.PublicServantUpdate{
				Department: req.Department,
			})

			return err
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("publicServantUpdate():", err)
			return
		}

		respondJSON(w, http.StatusOK, publicServantUpdateResponse{
			PublicServant: publicServant,
		})
	})
}

func (h *handler) publicServantDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			return q.PublicServantDelete(r.Context(), email)
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("publicServantDelete():", err)
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}
