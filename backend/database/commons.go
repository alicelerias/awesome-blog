package database

import "github.com/alicelerias/blog-golang/config"

func (s *PostgresDBRepository) GetLimit() string {
	return config.GetConfig().Limit
}
