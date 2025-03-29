package main

import "net/http"

func (app *Config) HomePage(responseWriter http.ResponseWriter, req *http.Request) {
	app.render(responseWriter, req, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(responseWriter http.ResponseWriter, req *http.Request) {
	app.render(responseWriter, req, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(responseWriter http.ResponseWriter, req *http.Request) {

}

func (app *Config) LogoutPage(responseWriter http.ResponseWriter, req *http.Request) {

}

func (app *Config) RegisterPage(responseWriter http.ResponseWriter, req *http.Request) {
	app.render(responseWriter, req, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(responseWriter http.ResponseWriter, req *http.Request) {

}

func (app *Config) ActivateAccountPage(responseWriter http.ResponseWriter, req *http.Request) {

}
