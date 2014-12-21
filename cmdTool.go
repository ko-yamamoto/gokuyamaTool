package main

import (
	"flag"
	"fmt"
	"github.com/nishikawasasaki/gokuyamaClient"
	"os"
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

func getKeysCommand(args []string) {

	tag := args[1]

	var gc gokuyamaClient.GokuyamaClient
	gc.Connect(*cmdHostname, *cmdPortNo)

	ret, err := gc.GetKeysByTag(tag)

	if err != nil {
		fmt.Errorf("error: %s\n", err)
	} else {
		fmt.Printf("result: %s\n", ret)
	}

}

func setCommand(args []string) {

	argsNum := len(flag.Args())

	if argsNum == 3 {
		// set without tags
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

	} else if argsNum == 4 {
		// set with tags
		key := args[1]
		value := args[2]
		tag := args[3]

		var gc gokuyamaClient.GokuyamaClient
		gc.Connect(*cmdHostname, *cmdPortNo)

		ret := gc.SetValueWithTag(key, value, tag)
		if ret == true {
			fmt.Println("registerd")
		} else {
			fmt.Errorf("error: %s\n", ret)
		}
	}
}

func showHelp() {
	msg := `Usage:
    set -- Set key and value.
    get -- Get value from the key
    tag -- Get keys from the tag`
	fmt.Println(msg)
}

func main() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		// show help
		showHelp()
		os.Exit(0)
	}

	commandName := &flag.Args()[0]
	// fmt.Println(commandName)

	switch {
	case *commandName == "set":
		setCommand(flag.Args())
	case *commandName == "get":
		getCommand(flag.Args())
	case *commandName == "tag":
		getKeysCommand(flag.Args())
	}

}
