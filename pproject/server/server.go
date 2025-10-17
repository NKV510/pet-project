package server

import (
	"net/http"

	handlers "github.com/NKV510/pproject/pkg/Handlers"
	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *handlers.HTTPHandlers
}

func NewHTTPServer(httpHandler *handlers.HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) HTTPServerStart() error {
	r := mux.NewRouter()
	r.Path("/worker").Methods("POST").HandlerFunc(s.httpHandlers.HandlersAddWorker)
	r.Path("/worker/{id}").Methods("PATCH").Queries("status", "true").HandlerFunc(s.httpHandlers.HandlersStartWork)
	r.Path("/worker/{id}").Methods("PATCH").Queries("status", "false").HandlerFunc(s.httpHandlers.HandlersEndWork)
	r.Path("/worker").Methods("GET").HandlerFunc(s.httpHandlers.HandlersGetAllWarkers)
	return http.ListenAndServe(":8080", r)
}
