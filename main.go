package main

import (
	"io"
	"os"
	"recserver/provider/config"
	"recserver/provider/http"
	public_routers "recserver/router"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiLogFile)

	env := config.ReadConfigigurationFile()
	port := env.GetString("PORT")
	engine := handlebars.New("./views/default", ".hbs")
	preFork := env.GetBool("FORK")

	server := http.NewServer(http.Server{
		Port:         port,
		RenderEngine: engine,
		PreFork:      preFork,
	})
	app := server.Setup()

	// ROUTERS CONFIGURATION
	routers := public_routers.NewRouters(app)
	routers.Setup()

	server.StartListening(app)
}
