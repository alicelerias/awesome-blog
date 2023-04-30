package controllers

import (
	"github.com/alicelerias/blog-golang/cache"
	"github.com/alicelerias/blog-golang/database"
)

type Server struct {
	repository database.Repository
	cache      cache.Repository
}

func NewServer(repository database.Repository, cache cache.Repository) *Server {
	return &Server{
		repository: repository,
		cache:      cache,
	}
}
