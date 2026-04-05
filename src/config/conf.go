package config

type Config struct {
	APP_PATH string `env:"APP_ROOT,default=/"`
	APP_PORT string `env:"APP_PORT,default=3000"`
}
