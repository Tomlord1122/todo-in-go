package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER" env:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE" env:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS" env:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// let viper to read the config file
	viper.AutomaticEnv() // read from environment variables
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config) // unmarshal the config into the struct
	return
}
