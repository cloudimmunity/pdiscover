package main

import (
	"log"

    "github.com/cloudimmunity/pdiscover"
)

func main() {
    exePath, err := pdiscover.GetOwnProcPath()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("exePath:",exePath)
}
