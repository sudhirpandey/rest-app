package main

import (
    "rest-app/app"
    "rest-app/config"
)

func main(){
    config := config.GetConfig()
    app := &app.App{}
    
    app.Initialize(config)
    app.Run(":8081")
}
