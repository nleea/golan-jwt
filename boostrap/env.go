package boostrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func LoadEnv() *Env {
	env := &Env{}

	viper.SetConfigFile("../.env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return env
}
