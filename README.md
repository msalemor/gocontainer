# gocontainer

A very simple GO singleton container.

## Getting the package

Get the package:

`go get -u github.com/msalemor/gocontainer"

## Package usage

### Define a service

```go
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
```

### Register and use the service in an application

```go
package main

import (
	"fmt"

	"github.com/msalemor/gocontainer/pkg"
	"github.com/msalemor/gocontainer/pkg/demo"
)

var container = pkg.AppContainer

func initializeServices() {
	var webService demo.APIServer = &demo.Server{Container: container}
	container.Register("web", webService)
}

func init() {
	initializeServices()
}

func main() {
	var webService = container.Get("web").(demo.APIServer)
	webService.Serve()
}
```
