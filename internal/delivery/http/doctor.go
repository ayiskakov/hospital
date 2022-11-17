package http

import (
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/khanfromasia/hospital/internal/entity"
	"github.com/khanfromasia/hospital/internal/storage"
)

type doctorCreateRequest struct {
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
	Doctor struct {
		Degree string `json:"degree"`
	} `json:"doctor"`
}

type doctorCreateResponse struct {
	Doctor entity.Doctor `json:"doctor"`
}

func (h *handler) doctorCreate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req doctorCreateRequest

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var doctor entity.Doctor

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			doctor.User, err = q.UserCreate(r.Context(), entity.User{
				Name:    req.User.Name,
				Surname: req.User.Surname,
				Email:   req.User.Email,
				Phone:   req.User.Phone,
				Salary:  req.User.Salary,
				Country: entity.Country{
					Cname: req.User.Country.Cname,
				},
			})

			doctor.Degree = req.Doctor.Degree

			doctor, err = q.DoctorCreate(r.Context(), doctor)

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusCreated, doctorCreateResponse{
			Doctor: doctor,
		})
	})
}

type doctorGetAllResponse struct {
	Doctors []entity.Doctor `json:"doctors"`
}

func (h *handler) doctorGetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctors []entity.Doctor

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			doctors, err = q.DoctorGetAll(r.Context())

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusOK, doctorGetAllResponse{
			Doctors: doctors,
		})
	})
}

type doctorUpdateRequest struct {
	Degree *string `json:"degree"`
}

type doctorUpdateResponse struct {
	Doctor entity.Doctor `json:"doctor"`
}

func (h *handler) doctorUpdate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req doctorUpdateRequest
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		if err := decodeJSONBody(w, r, &req); err != nil {
			respondError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var doctor entity.Doctor

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			doctor, err = q.DoctorUpdate(r.Context(), email, entity.DoctorUpdate{
				Degree: req.Degree,
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

		respondJSON(w, http.StatusOK, doctorUpdateResponse{
			Doctor: doctor,
		})
	})
}

func (h *handler) doctorDelete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := httprouter.ParamsFromContext(r.Context()).ByName("email")

		err := h.storage.ExecTX(r.Context(), pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}, func(q *storage.Queries) error {
			var err error

			err = q.DoctorDelete(r.Context(), email)

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
