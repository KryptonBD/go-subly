package main

import (
	"net/http"
)

func (app *Config) HomePage(responseWriter http.ResponseWriter, req *http.Request) {
	app.render(responseWriter, req, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(responseWriter http.ResponseWriter, req *http.Request) {
	app.render(responseWriter, req, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(responseWriter http.ResponseWriter, req *http.Request) {
	_ = app.Session.RenewToken(req.Context())

	err := req.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	email := req.Form.Get("email")
	password := req.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)

	if err != nil {
		app.Session.Put(req.Context(), "error", "Invalid Credentials")
		http.Redirect(responseWriter, req, "/login", http.StatusSeeOther)

		return
	}

	validPassword, err := user.PasswordMatches(password)

	if err != nil || !validPassword {
		app.Session.Put(req.Context(), "error", "Invalid Credentials")
		http.Redirect(responseWriter, req, "/login", http.StatusSeeOther)

		return
	}

	app.Session.Put(req.Context(), "userID", user.ID)
	app.Session.Put(req.Context(), "user", user)

	app.Session.Put(req.Context(), "flash", "Login successful!")

	http.Redirect(responseWriter, req, "/", http.StatusSeeOther)
}

func (app *Config) LogoutPage(responseWriter http.ResponseWriter, req *http.Request) {
	_ = app.Session.Destroy(req.Context())
	_ = app.Session.RenewToken(req.Context())

	http.Redirect(responseWriter, req, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(responseWriter http.ResponseWriter, req *http.Request) {
	app.render(responseWriter, req, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(responseWriter http.ResponseWriter, req *http.Request) {

}

func (app *Config) ActivateAccountPage(responseWriter http.ResponseWriter, req *http.Request) {

}
