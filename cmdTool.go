package main

import (
	"flag"
	"fmt"
	"github.com/nishikawasasaki/gokuyamaClient"
)

var (
	cmdHostname = flag.String("h", "localhost", "hostname of okuyama master node")
	cmdPortNo   = flag.Int("p", 8888, "port number of okuyama master node")
)

func getCommand(args []string) {

	key := args[1]

	var gc gokuyamaClient.GokuyamaClient
	gc.Connect(*cmdHostname, *cmdPortNo)

	ret, err := gc.GetValue(key)

	if err != nil {
		fmt.Errorf("error: %s\n", err)
	} else {
		fmt.Printf("result: %s\n", ret)
	}

}

func setCommand(args []string) {

	key := args[1]
	value := args[2]

	var gc gokuyamaClient.GokuyamaClient
	gc.Connect(*cmdHostname, *cmdPortNo)

	ret := gc.SetValue(key, value)
	if ret == true {
		fmt.Println("registerd")
	} else {
		fmt.Errorf("error: %s\n", ret)
	}

}

func main() {

	flag.Parse()

	commandName := &flag.Args()[0]
	// fmt.Println(commandName)

	// if *key == "" && *value == "" {
	// if *key == "" {
	// fmt.Println("Option -k is required")
	// os.Exit(1)
	// }

	switch {
	case *commandName == "set":
		setCommand(flag.Args())
	case *commandName == "get":
		getCommand(flag.Args())
	}

}
