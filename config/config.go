package config

import (
	"os"
	"log"
	"strings"
	"github.com/spf13/viper"
)


type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Name     string
	Charset  string
}


// Use the viper to load the environent variable if the Environment is not dev. Return the loadéd config
// either from environmnet or from config file as config struct. 
// TODO viper also gives methods for unmarshal the config into structs

func GetConfig() *Config {
    if os.Getenv("ENVIRONMENT") == "DEV" {        
		viper.SetConfigName("config")  
		viper.SetConfigType("yaml") 
	    viper.AddConfigPath("./config")  
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
   } else {  
	   viper.SetEnvPrefix("DB")
	   viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	   viper.AutomaticEnv()

	   username := viper.GetString("username")
	   println(viper.GetString(username))
	}

	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: viper.GetString("username"),
			Password: viper.GetString("password"),
			Host:     viper.GetString("host"),
			Name:     viper.GetString("name"),
			Charset:  "utf8",
		},
	}
}