package html

import (
	"fmt"
	"net/http"

	"demo.com/demo"
)

type CarHandler struct {
	Car demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<font color='green'><b>Make: </b></font>  <i>%v<i><br><font color='green'><b>Model: </b></font>  <i>%v<i><br><font color='green'><b>Year: </b></font>  <i>%v<i><br>", h.Car.Make, h.Car.Model, h.Car.Year)
}

type PersonHandler struct {
	Person demo.Person
}

func (p *PersonHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<font color='green'><b>Name: </b></font>  <i>%v %v<i><br>", p.Person.FirstName, p.Person.LastName)
}
