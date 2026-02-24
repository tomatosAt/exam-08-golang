package config

type Config struct {
	App    appConfig
	Server serverCfg
	DBMain MariaDBConfig
}

type appConfig struct {
	Name       string
	Version    string
	Mode       string
	PrefixPath string
	StorageDir string
	Url        string
}
