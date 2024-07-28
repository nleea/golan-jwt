package boostrap

import "database/sql"

type Application struct {
	DB  *sql.DB
	ENV *Env
}

func InitApp() Application {

	app := &Application{}

	app.ENV = LoadEnv()
	app.DB = DatabaseConnect(app.ENV)

	return *app
}

func (app *Application) DatabaseDisconnect() {
	DatabaseDisconnect(app.DB)
}
