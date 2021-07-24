package router

import "github.com/kataras/iris/v12"

func Set(app *iris.Application) {
	r := app.Party("/api")

	setRoom(r)
}
