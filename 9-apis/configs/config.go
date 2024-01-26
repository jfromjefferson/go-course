package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("H256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
