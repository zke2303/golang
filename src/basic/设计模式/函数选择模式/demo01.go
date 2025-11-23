package main

type DataSource struct {
	host     string
	port     string
	username string
	password string
	database string
}

type Option func(*DataSource)

func WithHost(host string) Option {
	return func(s *DataSource) {
		s.host = host
	}
}

func WithPort(port string) Option {
	return func(s *DataSource) {
		s.port = port
	}
}
func WithUsername(username string) Option {
	return func(s *DataSource) {
		s.username = username
	}
}

func WithPassword(password string) Option {
	return func(s *DataSource) {
		s.password = password
	}
}

func WithDatabase(database string) Option {
	return func(s *DataSource) {
		s.database = database
	}
}
