package main // viewData represents the root model used to dynamically update the app

import (
	"html/template"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	template.Must(templates.ParseGlob("internal/*/*/*"))
}

func main() {
	if len(logFilePath) > 1 {
		f := setupLogging()
		defer f.Close()
	}

	ctx, srv := bolt()

	log.Println("Waiting for connections @ http://localhost" + srv.Addr)

	<-ctx.Done()
}
