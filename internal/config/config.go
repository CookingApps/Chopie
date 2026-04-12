package config

import "github.com/spf13/viper"

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() (config Config, err error) {
	// 1. Tell Viper where .env is (current directory)
	viper.AddConfigPath(".")
	// 2. Tell Viper the file name and type
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// 3. Load the file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// 4. & 5. Unmarshal into the Config struct and return it
	err = viper.Unmarshal(&config)
	return
}
