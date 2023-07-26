package postgres

type ClientConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
