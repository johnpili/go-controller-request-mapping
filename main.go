package main

import (
	"flag"
	"github.com/go-zoo/bone"
	"github.com/johnpili/go-controller-request-mapping/controllers"
	"github.com/johnpili/go-controller-request-mapping/models"
	"github.com/psi-incontrol/go-sprocket/sprocket"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	configuration models.Config
)

func main() {
	pid := os.Getpid()
	err := ioutil.WriteFile("application.pid", []byte(strconv.Itoa(pid)), 0666)
	if err != nil {
		log.Fatal(err)
	}

	var configLocation string
	flag.StringVar(&configLocation, "config", "config.yml", "Set the location of configuration file")
	flag.Parse()

	err = sprocket.LoadYAML(configLocation, &configuration)
	if err != nil {
		log.Fatal(err)
	}

	decimal.MarshalJSONWithoutQuotes = true

	controllersHub := controllers.New()
	router := bone.New()
	//router.Get("/static/", staticFileServer)
	controllersHub.BindRequestMapping(router)

	// CODE FROM https://medium.com/@mossila/running-go-behind-iis-ce1a610116df
	port := strconv.Itoa(configuration.HTTP.Port)
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}

	//csrfProtection := csrf.Protect(
	//	[]byte(configuration.System.CSRFKey),
	//	csrf.Secure(false),
	//)

	httpServer := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  900 * time.Second,
		WriteTimeout: 900 * time.Second,
		Handler: router,
	}

	if configuration.HTTP.IsTLS {
		log.Printf("Server running at https://localhost:%s/\n", port)
		log.Fatal(httpServer.ListenAndServeTLS(configuration.HTTP.ServerCert, configuration.HTTP.ServerKey))
		return
	}
	log.Printf("Server running at http://localhost:%s/\n", port)
	//log.Fatal(http.ListenAndServe(":"+port, router)) // Start HTTP Server

	log.Fatal(httpServer.ListenAndServe())
}
