package main

import (
	"context"
	"fmt"
	"github.com/totalsecond/quotesapi/db"
	"github.com/totalsecond/quotesapi/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var dbfile = "quotes.json"

func main() {
	fmt.Println("Starting the quotes internet machine")

	db.Load(dbfile)
	router.SetupRouter()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Received shutdown")
	db.Save(dbfile)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server shutdown: ", err)
	}
}
