package config

import (
	"log"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/spf13/viper"
)

var (
	JWT_SECRET string
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
}

type CldConfig struct {
	CLDNAME     string
	CLDKEY      string
	CLDSECRET   string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	cld := CldConfig{}
	isRead := true

	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWTSECRET"); found {
		JWT_SECRET = val
		isRead = false
	}

	if val, found := os.LookupEnv("CLDNAME"); found {
		cld.CLDNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLDKEY"); found {
		cld.CLDKEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLDSECRET"); found {
		cld.CLDSECRET = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		JWT_SECRET = viper.GetString("JWTSECRET")
		cld.CLDNAME = viper.Get("CLDNAME").(string)
		cld.CLDKEY = viper.Get("CLDKEY").(string)
		cld.CLDSECRET = viper.Get("CLDSECRET").(string)
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
	}

	return &app
}

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(cld.CLDNAME, cld.CLDKEY, cld.CLDSECRET)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
