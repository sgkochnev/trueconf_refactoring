package v1

import (
	"net/http"
	"refactoring/internl/dto"
	repoerror "refactoring/internl/repository/repo_error"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type userController struct {
	UserUseCase
}

func NewUserController(uCase UserUseCase) *userController {
	return &userController{
		UserUseCase: uCase,
	}
}

func (c *userController) SearchUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserUseCase.SearchUsers()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		_ = render.Render(w, r, ErrInternal())
	}
	render.JSON(w, r, users)
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {

	request := dto.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id, err := c.UserUseCase.CreateUser(&request)
	if err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (c *userController) GetUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	user, err := c.UserUseCase.GetUser(id)
	if err != nil {
		render.Render(w, r, ErrInternal())
	}

	render.JSON(w, r, user)
}

func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	request := dto.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	err := c.UserUseCase.UpdateUser(id, &request)
	if err != nil {
		var ErrResp render.Renderer
		if err == repoerror.ErrUserNotFound {
			ErrResp = ErrInvalidRequest(err)
		} else {
			ErrResp = ErrInternal()
		}
		_ = render.Render(w, r, ErrResp)
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	err := c.UserUseCase.DeleteUser(id)
	if err != nil {
		var ErrResp render.Renderer
		if err == repoerror.ErrUserNotFound {
			ErrResp = ErrInvalidRequest(err)
		} else {
			ErrResp = ErrInternal()
		}
		_ = render.Render(w, r, ErrResp)
		return
	}

	render.Status(r, http.StatusNoContent)
}
