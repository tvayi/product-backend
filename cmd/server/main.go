package main

import (
    "fmt"
    "product-backend/pkg/api"
    "product-backend/pkg/db"
    "log"
    "time"
    "net/http"
    "os"
)

func main() {
    log.Print("server has started")

    time.Sleep(10 * time.Second)

    pgdb, err := db.StartDB()
    if err != nil {
        log.Printf("error starting the database %v", err)
    }

    router := api.StartAPI(pgdb)

    port := os.Getenv("PORT")

    err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
    if err != nil {
        log.Printf("error from router %v\n", err)
    }
}
