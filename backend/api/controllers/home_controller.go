package controllers

import (
	"net/http"

	"github.com/alicelerias/blog-golang/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hello, World")
}
