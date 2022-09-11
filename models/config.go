package models

type Config struct {
	WebServer Server
	Database  Database
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
