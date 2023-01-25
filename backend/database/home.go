package database

func (s *PostgresDBRepository) GetHome() error {
	return s.db.DB().Ping()
}
