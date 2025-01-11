package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

//create a config structure  for store the config data...

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

//create a function for checking u the config files

func MustLoad() *Config {
	var configPath string
	//using this to get the env path..
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		//if we are passing any flags while running the code then this code find that and parse it
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()

		configPath = *flags

		//handel the error if we are not passing any path of config file
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	//here we check if there are any error while finding the path then we throw the error
	// _,err:=os.Stat(configPath)
	//if os.IsNotExist(err){
	// 		log.Fatalf("config file does not exit :%s", configPath)

	// }
	//the meaning is same in one line

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exit :%s", configPath)
	}

	var cfg Config
	//if everything is fine then by using cleanenv we read the confog file or the env file...
	err := cleanenv.ReadConfig(configPath, &cfg)

	//if there are any error while reading the file it throw error..
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	//return the config address...
	return &cfg

}
