package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/jfromjefferson/gi-course-9/internal/dto"
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	"github.com/jfromjefferson/gi-course-9/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	userDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIs int
}

func NewUserHandler(db database.UserInterface, JWT *jwtauth.JWTAuth, JwtExpiresIs int) *UserHandler {
	return &UserHandler{
		userDB:       db,
		Jwt:          JWT,
		JwtExpiresIs: JwtExpiresIs,
	}
}

// GetJWT godoc
// @Summary Get JWT
// @Description Get JWT...
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.JWTInput true "user credentials"
// @Success 200 {object} dto.JWTOutput
// @Failure 404
// @Failure 500 {object} Error
// @Router /auth [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var JwtDTO dto.JWTInput

	err := json.NewDecoder(r.Body).Decode(&JwtDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userDB.FindByEmail(JwtDTO.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(JwtDTO.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIs)).Unix(),
	})

	accessToken := dto.JWTOutput{
		AccessToken: token,
	}

	err = json.NewEncoder(w).Encode(accessToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Create user godoc
// @Summary Create user
// @Description Create user...
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.FirstName, userDTO.LastName, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !user.ValidatePassword(userDTO.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.userDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
