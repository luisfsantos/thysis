package daemon

import (
	"github.com/luisfsantos/thysis/db"
	"os"
	"os/signal"
	"syscall"
	"log"
	"github.com/luisfsantos/thysis/server"
	"net"
	"github.com/luisfsantos/thysis/model"
)

type Configuration struct {
	BindingAddress string

	DBconfig db.Configuration
	UI       server.Configuration
}

func Run(configuration *Configuration) error {

	log.Println("Running Daemon...")
	log.Printf("Starting, HTTP on: %s\n", configuration.BindingAddress)

	//db, err := db.InitDb(cfg.Db)
	//if err != nil {
	//	log.Printf("Error initializing database: %v\n", err)
	//	return err
	//}
	//
	//m := model.New(db)
	db, err := db.InitDB(configuration.DBconfig)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		return err
	}
	l, err := net.Listen("tcp", configuration.BindingAddress)
	if err != nil {
		log.Printf("Error creating listener: %v\n", err)
		return err
	}

	server.Start(configuration.UI, &model.Model{DB:db}, l)
	waitForSignal()
	return nil

}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
