package plaintext

import (
	"fmt"
	"net/http"

	"demo.com/demo"
)

type CarHandler struct {
	Car demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", h.Car.Make, h.Car.Model, h.Car.Year)
}

type PersonHandler struct {
	Person demo.Person
}

func (p *PersonHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Name: %v %v", p.Person.FirstName, p.Person.LastName)
}
