package main

import (
    "postwoman/utils"
    "postwoman/handlers"
)

func main() {
	
    var env = utils.GetEnv()
    e := handlers.ConfigGlobalHandler()

    e = handlers.TemplateHandler()
    e = handlers.UserHandler()
    e = handlers.RequestHandler()

    e.Logger.Fatal(e.Start(env["GO_PORT"]))
}
