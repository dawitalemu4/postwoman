package main

import (
    "postwoman/utils"
    "postwoman/routes"
)

func main() {

    var env = utils.GetEnv()
    e := routes.ConfigGlobalRoutes()

    e = routes.TemplateRoutes()
    e = routes.UserRoutes()
    e = routes.RequestRoutes()

    e.Logger.Fatal(e.Start(env["GO_PORT"]))
}
