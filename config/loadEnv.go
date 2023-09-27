package config

import "github.com/spf13/viper"

type env struct {
	// DataBase Setup
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUserame  string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBENGINE   string `mapstructure:"DB_ENGINE"`
	// Redis Setup
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	// Server Setting
	AppUrl  string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"APP_PORT"`
	Issuer  string `mapstructure:"ISSUER"`
	// jwt secret
	JBT    string `mapstructure:"JWT_BUFFER_TIME"`
	JET    string `mapstructure:"JWT_EXPIRE_TIME"`
	Secret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (config env, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// handle null
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
