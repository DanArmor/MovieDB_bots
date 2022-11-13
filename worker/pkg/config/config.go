package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	SqlUrl     string `mapstructure:"SqlUrl"`
	AdminPass  string `mapstructure:"ADMIN_PASS"`
	CertPath   string `mapstructure:"CERT_PATH"`
	KeyPath    string `mapstructure:"KEY_PATH"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	PathServer string `mapstructure:"PATH_SERVER"`
	ServerURL  string
	SqlUser    string `mapstructure:"SQL_USER"`
	SqlPass    string `mapstructure:"SQL_PASS"`
	Domain     string `mapstructure:"DOMAIN"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
