package demo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/msalemor/gocontainer/pkg"
	"github.com/sirupsen/logrus"
)

type APIServer interface {
	Serve()
}

type Server struct {
	Container *pkg.Container
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}

func (s *Server) Serve() {
	http.HandleFunc("/api/hello", helloHandler)
	log.Println("Server started on port http://localhost:8180")
	logrus.Info(http.ListenAndServe(":8180", nil))
}
