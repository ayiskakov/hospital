package http

import (
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/backend/internal/entity"
	"github.com/khanfromasia/hospital/backend/internal/storage"
)

type userCreateRequest struct {
	User struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Salary  int64  `json:"salary"`
		Country struct {
			Cname string `json:"cname"`
		} `json:"country"`
	} `json:"user"`
}

type userCreateResponse struct {
	User entity.User `json:"user"`
}

func (h *handler) userCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req userCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var user entity.User

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			user, err = q.UserCreate(r.Context(), entity.User{
				Name:    req.User.Name,
				Surname: req.User.Surname,
				Email:   req.User.Email,
				Phone:   req.User.Phone,
				Salary:  req.User.Salary,
				Country: entity.Country{
					Cname: req.User.Country.Cname,
				},
			})

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusCreated, userCreateResponse{
			User: user,
		})
	})
}

type userUpdateRequest struct {
	Name    *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Salary  *int64  `json:"salary,omitempty"`
	Country struct {
		Cname *string `json:"cname,omitempty"`
	}
}

type userUpdateResponse struct {
	User entity.User `json:"user"`
}

func (h *handler) userUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req userUpdateRequest
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var user entity.User

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			user, err = q.UserUpdate(r.Context(), email, entity.UserUpdate{
				Name:    req.Name,
				Surname: req.Surname,
				Phone:   req.Phone,
				Salary:  req.Salary,
				Cname:   req.Country.Cname,
			})

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, userUpdateResponse{
			User: user,
		})
	})
}

func (h *handler) userDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			err = q.UserDelete(r.Context(), email)

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, nil)
	})
}

type userGetAllResponse struct {
	Users []entity.User `json:"users"`
}

func (h *handler) userGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var users []entity.User

		err := h.storage.Transaction(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			users, err = q.UserGetAll(r.Context())

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			log.Println("userGetAll():", err)
			return
		}

		respondJSON(w, http.StatusOK, userGetAllResponse{
			Users: users,
		})
	})
}
