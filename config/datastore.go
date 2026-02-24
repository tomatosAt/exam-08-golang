package config

type MariaDBConfig struct {
	Host      string
	Port      string
	User      string
	Password  string
	Database  string
	Migration bool
}
