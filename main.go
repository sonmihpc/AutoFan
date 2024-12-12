package main

import (
	"github.com/coreos/go-systemd/v22/daemon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := readConf()
	controller := NewControllerFromConfig(config)

	controller.Run()

	if _, err := daemon.SdNotify(false, "READY=1"); err != nil {
		log.Print("system notification supported, but failed.")
	}

	// prohibit the terminal of main process
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGILL, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGHUP, syscall.SIGILL, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("autofan exit.")
				ch <- s
			default:
				log.Printf("received other singnal: %s, sonmid exit.", s)
			}
		}
	}()
	<-ch
}
