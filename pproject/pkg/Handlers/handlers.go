package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NKV510/pproject/pkg/worker"
	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	worker worker.List
}

func NewHTTPHandlers(worker *worker.List) *HTTPHandlers {
	return &HTTPHandlers{
		worker: *worker,
	}
}

/*
end_point -- /work
method -- POST
Request -- JSON body
*/
func (h *HTTPHandlers) HandlersAddWorker(w http.ResponseWriter, r *http.Request) {
	var worker worker.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	id := h.worker.AddWorker(worker.Name, worker.SecondName)
	fmt.Println(id)
	str := "Вам присвоен ID:" + id
	retStr, err := json.Marshal(str)
	if err != nil {
		return
	}
	if _, err := w.Write(retStr); err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

}

/*
end_point -- /work/{id}?status={bool}
method -- PATCH
-----
response body:
JSON
*/
func (h *HTTPHandlers) HandlersStartWork(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := h.worker.StartWork(id)
	retUser, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	if _, err := w.Write(retUser); err != nil {
		panic(err)
	}

}

/*
end_point -- /work/{id}?status={bool}
method -- PATCH
-----
response body:
JSON
*/
func (h *HTTPHandlers) HandlersEndWork(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user := h.worker.EndWork(id)
	retUser, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	if _, err := w.Write(retUser); err != nil {
		panic(err)
	}
}

/*
end_point -- /work
method -- GET
------
*/
func (h *HTTPHandlers) HandlersGetAllWarkers(w http.ResponseWriter, r *http.Request) {
	var slice any
	slice = h.worker.GetAllWorkers()
	retSlice, err := json.MarshalIndent(slice, "", "	")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if _, err := w.Write(retSlice); err != nil {
		panic(err.Error())
	}
}

func (h *HTTPHandlers) HandlersDeletWarker(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	h.worker.DeleteWorker(id)
}
