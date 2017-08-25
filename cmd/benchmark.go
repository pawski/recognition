package main

import "time"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {

	url := os.Args[1]

	timer := time.NewTimer(time.Second * 1)
	hit := 0

	go gracefulShutdown()
	ticker := time.NewTicker(time.Millisecond * 1)
	go func() {
		for t := range ticker.C {
			log.Println("Start at", t)
			getUrl(url)
			hit++
		}
	}()

	<-timer.C
	log.Printf("Did %d hits", hit)
}

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		// clean up here
		os.Exit(0)
	}()
}

func getUrl(url string)  {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	//writeToFile(err, body)

	log.Printf("%d bytes", len(body))
	log.Print(resp.Status)
}

func writeToFile(err error, body []byte) {
	t := time.Now()
	err = ioutil.WriteFile("./cache/"+strconv.FormatInt(t.Unix(), 10)+".html", body, 0644)
	check(err)
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}