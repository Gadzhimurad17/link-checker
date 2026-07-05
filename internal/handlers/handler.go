package handlers

import (
	"net/http"
	 "github.com/joho/godotenv"
)

type Server struct {
	mux    *http.ServeMux
	server http.Server
}
func CreateServe(){
	http.ListenAndServe(":8080",nil)
}
func HandleRequest() {
	http.HandlerFunc()
}
func main() {

}
