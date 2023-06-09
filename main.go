package main

import (
	"flag"
	"fmt"
	"gostore/command"
	"gostore/logger"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	ch "gostore/commandhandler"
)

const (
	argsMsg    = "Reading Command Line Arguments..."
	usageMsg   = "A port to run the gostore on"
	helpMsg    = "A port was not specified"
	startedMsg = "Started gostore application on port: %d\n"
)

func main() {

	defer logger.Close()

	logger.InfoLogger.Println(argsMsg)
	port := flag.Int("port", 8080, usageMsg)
	flag.Parse()

	if port := isFlagPassed("port"); !port {
		fmt.Println(helpMsg)
		logger.FatalLogger.Println(helpMsg)
		os.Exit(-1)
	}

	fmt.Printf(startedMsg, *port)
	logger.InfoLogger.Printf(startedMsg, *port)

	portStr := "localhost:" + strconv.Itoa(*port)

	http.HandleFunc("/", genericHandler)
	if err := http.ListenAndServe(portStr, nil); err != nil {
		logger.FatalLogger.Println("failed to start server")
	}

	logger.InfoLogger.Println("Ended gostore Application")
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/test") {
		ch.HandleCommand(w, command.Command{Verb: command.TestCmd})
	} else if r.Method == http.MethodGet {
		key := path.Base(r.URL.String())
		fmt.Printf("\nGET: %v", key)
		ch.HandleCommand(w, command.Command{Verb: command.GetCmd, Args: command.Arguments{Key: key}})
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("\nPOST: %s = %s", key, value)
			ch.HandleCommand(w, command.Command{Verb: command.PutCmd, Args: command.Arguments{Key: key, Value: value[0]}})
		}
	} else if r.Method == http.MethodDelete {
		key := path.Base(r.URL.String())
		fmt.Printf("\nDEL: %v", key)
		ch.HandleCommand(w, command.Command{Verb: command.DelCmd, Args: command.Arguments{Key: key}})
	}
}
