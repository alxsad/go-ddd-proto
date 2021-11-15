package web

import (
	"app/app"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

type WebServer struct {
	app app.Application
}

func NewWebServer(port int, timeout time.Duration, app app.Application) *http.Server {
	var webServer = &WebServer{app: app}

	var server = http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     webServer.newRouter(),
		ReadTimeout: timeout,
	}

	return &server

}

func (self *WebServer) newRouter() chi.Router {
	var r = chi.NewRouter()

	r.Get("/users", self.ListUsers)
	r.Post("/users", self.CreateUser)

	return r
}
