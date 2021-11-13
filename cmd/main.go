package main

import (
	"log"

	"github.com/User0608/zeus_project_api/cmd/auth"
	"github.com/User0608/zeus_project_api/cmd/injectors"
	"github.com/User0608/zeus_project_api/cmd/router"
	"github.com/User0608/zeus_project_api/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf, err := configs.LoadServiceConfigs("service_config.json")
	if err != nil {
		log.Fatalln("No se cargo las configuraciones de servicio:", err.Error())
	}
	log.Println("Server configs OK!")
	if err := auth.LoadFiles(conf.Certificates.Private, conf.Certificates.Public); err != nil {
		log.Fatalln("No se cargaron los certificados,", err.Error())
	}
	log.Println("Certificados OK!")
	if err := injectors.LoadConfig("db_config.json"); err != nil {
		log.Fatalln("Error base de datos,", err.Error())
	}
	log.Println("Base de datos OK!")

	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.Cors.AllowOrigins,
		AllowMethods: conf.Cors.AllowMethods,
	}))
	server.HideBanner = true
	router.Upgrade(server)
	if err := server.Start(conf.Address); err != nil {
		log.Fatal()
	}
	log.Println("Server detenido!")
}
