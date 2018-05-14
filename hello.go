package main

import "fmt"
import (
    log "github.com/sirupsen/logrus"
)

func main(){
    fmt.Println("Hello, Golang!")
    log.WithFields(log.Fields{
        "animal": "walrus",
    }).Info("A walrus appears")
}