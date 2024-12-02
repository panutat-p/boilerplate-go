package config

type Config struct {
	AppName string `env:"APP_NAME" validate:"required"`
	Version string `env:"VERSION" validate:"required"`
	Port    string `env:"PORT" validate:"required"`
}
