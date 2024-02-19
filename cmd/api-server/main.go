package main

import (
	"ikiler-dosya/config"

	"ikiler-dosya/pkg/helpers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	serverOrigin := helpers.SafeGetEnv("SERVER_ORIGIN")
	audience := helpers.SafeGetEnv("AUTH0_AUDIENCE")
	auth0Domain := helpers.SafeGetEnv("AUTH0_DOMAIN")

	config := Config{
		ServerOrigin:  serverOrigin,
		SecureOptions: config.SecureOptions(),
		Audience:      audience,
		Auth0Domain:   auth0Domain,
	}

	app := App{Config: config}

	app.RunServer()
}
