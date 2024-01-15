package main

import (
	"file_transfer_protocol/file_components"
	"flag"
	"fmt"
	"net"
	"os"
)

// This functions handles logic behind args or flag of cli commands.
func argHandler(filepath *string, useQuit *bool, useViewFile *string, usePreview *bool, currentIp *bool) {

	// if flag.NArg() == 0 || flag.Arg(0) != "fsa" {
	// 	fmt.Println("correct command required")
	// 	os.Exit(1)
	// }
	if *useQuit {

		quitFlag(useQuit)
	}

	if *filepath != "" {

		file_components.ReadFileFromPath(filepath)
	}

	if *currentIp {

		displayAddress()
		os.Exit(0)
	}

	if *useViewFile == "" {
		flag.Usage()
		os.Exit(1)

	} else {
		file_components.MarkDownRun(*useViewFile, *usePreview)
	}

}

// this function when called exits the cli
func quitFlag(useQuit *bool) {
	if *useQuit {

		// logic if user enter "-q or --q arguments"
		os.Exit(0)
	}
}

func displayAddress() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		return
	}

	for _, iface := range interfaces {

		// Get the addresses associated with the current interface
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error getting addresses for interface", iface.Name, ":", err)
			continue
		}

		// Print the IP addresses associated with the current interface
		fmt.Println("IP Addresses for interface", iface.Name+":")
		for _, addr := range addrs {
			fmt.Println(addr)
		}
	}
	return

}
