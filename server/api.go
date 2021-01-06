package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	apiServer := &api{}

	router := mux.NewRouter()
	router.HandleFunc("/", apiServer.web).Methods(http.MethodGet)
	router.HandleFunc("/api/start", apiServer.start).Methods(http.MethodPost)
	router.HandleFunc("/api/timer", apiServer.timer).Methods(http.MethodPost)

	apiServer.router = router
	return apiServer
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) web(w http.ResponseWriter, r *http.Request) {}
func (a *api) start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
func (a *api) timer(w http.ResponseWriter, r *http.Request) {}

// in order to support the best accuracy it is expected to work in fractions
// so having 1 and 1/12th of f-stop would mean to have
//   fstops: 13, step: 12
// which will be more accuratelly transmitted than 1,083333333
// delay will give some time in seconds before starting the process
// in case it needs some preparations (like closing a laptops's lid)
type FStopTimer struct {
	Fstops int `json:"fstops,float"`
	Step   int `json:"step"`
	Delay  int `json:"delay,omitempty,int"`
}
