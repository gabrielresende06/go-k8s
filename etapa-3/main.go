package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

func Home(w http.ResponseWriter, r *http.Request)  {
	greating := greating("Code.education Rocks!")
	data := []byte(greating)
	w.Write(data)
}

func greating(term string) string {
	return fmt.Sprintf("<b>%s</b>", term)
}