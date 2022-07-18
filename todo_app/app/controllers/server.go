package controllers

import (
	"hello-module/todo_app/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
