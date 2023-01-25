package controllers

import "github.com/alicelerias/blog-golang/database"

type Server struct {
	repository database.Repository
}

func NewServer(repository database.Repository) *Server {
	return &Server{repository: repository}
}
