package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"subly/data"
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
		msg := Message{
			To:      email,
			Subject: "Failed log in attempt",
			Data:    "Invalid login attempt",
		}
		app.sendEmail(msg)

		app.Session.Put(req.Context(), "error", "Invalid Credentials")
		http.Redirect(responseWriter, req, "/login", http.StatusSeeOther)

		return
	}

	app.Session.Put(req.Context(), "userId", user.ID)
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
	err := req.ParseForm()

	if err != nil {
		app.ErrorLog.Println(err)
	}

	u := data.User{
		Email:     req.Form.Get("email"),
		FirstName: req.Form.Get("first-name"),
		LastName:  req.Form.Get("last-name"),
		Password:  req.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
	}

	_, err = u.Insert(u)

	if err != nil {
		app.Session.Put(req.Context(), "error", "Unable to create user.")
		http.Redirect(responseWriter, req, "/register", http.StatusSeeOther)
		return
	}

	FE_URL := os.Getenv("FE_URL")
	url := fmt.Sprintf("http://%s/activate?email=%s", FE_URL, u.Email)
	signedURL := GenerateTokenFromString(url)
	app.InfoLog.Println(signedURL)

	msg := Message{
		To:       u.Email,
		Subject:  "Activate your account",
		Template: "confirmation-email",
		Data:     template.HTML(signedURL),
	}

	app.sendEmail(msg)
	app.Session.Put(req.Context(), "flash", "confirmation email sent, please check your email")

	http.Redirect(responseWriter, req, "/login", http.StatusSeeOther)
}

func (app *Config) ActivateAccountPage(responseWriter http.ResponseWriter, req *http.Request) {
	url := req.RequestURI
	FE_URL := os.Getenv("FE_URL")
	tokenURL := fmt.Sprintf("http://%s%s", FE_URL, url)
	okay := VerifyToken(tokenURL)

	if !okay {
		app.Session.Put(req.Context(), "error", "Invalid Token")
		http.Redirect(responseWriter, req, "/", http.StatusSeeOther)
		return
	}

	u, err := app.Models.User.GetByEmail(req.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(req.Context(), "error", "No User found")
		http.Redirect(responseWriter, req, "/", http.StatusSeeOther)
		return
	}

	u.Active = 1
	err = u.Update()

	if err != nil {
		app.Session.Put(req.Context(), "error", "Unable to activate User")
		http.Redirect(responseWriter, req, "/", http.StatusSeeOther)
		return
	}

	app.Session.Put(req.Context(), "flash", "Account activated")
	http.Redirect(responseWriter, req, "/login", http.StatusSeeOther)

}
