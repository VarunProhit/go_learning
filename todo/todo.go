package main

import (
	// "fmt"
	// "net/htrtp"
    // "log"
	"example.com/server"
	// "github.com/gorilla/mux"
)


func main(){
  app := server.App{}
  app.Port =":9003"
  app.Initialize()
  app.Run()
}