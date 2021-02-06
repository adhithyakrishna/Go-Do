package controllers

import (
	"Simple-Form-Submission/models"
	"fmt"

	"net/http"

	"Simple-Form-Submission/views"

	"github.com/gorilla/schema"
)

type Users struct {
	LoginView *views.View
	us        *models.UserInfoService
}

func NewUsers(us *models.UserInfoService) *Users {
	return &Users{
		LoginView: views.NewView("bootstrap", "users/signup"),
		us:        us,
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.LoginView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.Userinfo{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
