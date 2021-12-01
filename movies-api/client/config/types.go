package config

type Server struct {
	Port               int
	LogFilePath        string `toml:"log_path"`
	DefaultHttpTimeout int    `toml:"http_timeout"`
}

type DB struct {
	DSN         string `toml:"dsn"`
	MaxConn     int    `toml:"max_conn"`
	MaxIdleConn int    `toml:"max_idle_conn"`
}
