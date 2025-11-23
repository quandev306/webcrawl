package bootstrap

import "github.com/quandev306/webcrawl/config"

type Application struct {
	Env   *Env
	Mongo config.Client
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
