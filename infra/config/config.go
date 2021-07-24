package config

// C
var C = &Config{}

// Config -
type Config struct {
	// Log - logger set
	Log Log
}

// Log -
type Log struct {
	Level string
}
