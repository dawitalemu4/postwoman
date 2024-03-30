package handlers

import "postwoman/utils"

var db = utils.DB()
var env = utils.GetEnv()

type jsonMessage struct {
    Key      string    `json:"key"`
    Value    string    `json:"value"`
}

func errorJSON(key string, value string) jsonMessage {
    return jsonMessage{Key: key, Value: value}
}
