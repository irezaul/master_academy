package main

import (
	"fmt"
	"net/http"
)

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
func main() {

	//var name datatype  ( eta variable tori korar signeture)
	//var name string  (eta uporer signeture onujayei variable toiri kora hoiyece)
	//var handler func(ResponseWriter, *Request) // eta 1ta datatype (func(ResponseWriter, *Request))

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8888", nil)
}

// JOKHON E AMRA func type er signeture dekbho tokon amar ai rokom syntex toiri korovo)

func home(w http.ResponseWriter, r *http.Request) { // var home string er moto
	fmt.Fprintf(w, "Welcome To my first webpage")

}

// about
func about(w http.ResponseWriter, r *http.Request) { // var home string er moto
	fmt.Fprintf(w, "Welcome To my about page")

}

// contact
func contact(w http.ResponseWriter, r *http.Request) { // var home string er moto
	fmt.Fprintf(w, "Welcome To my Contact page")

}
