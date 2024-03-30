package main

import (
    "postwoman/utils"
    "postwoman/routes"
)

func main() {

    var env = utils.GetEnv()
    e := routes.ConfigGlobalHandler()

    e = routes.TemplateHandler()
    e = routes.UserHandler()
    e = routes.RequestHandler()

    e.Logger.Fatal(e.Start(env["GO_PORT"]))
}
