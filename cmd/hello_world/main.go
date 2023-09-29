package main

import (
	"flag"
	"log"
	"runtime"
	"flags-cli-template/internal/name_package"
	"flags-cli-template/pkg/age_package"
)

func main() {
	var name = flag.String("name", "", "name to be outputted")
	var age= flag.Int("age", 0, "age to be outputted")
	var cpuNum = flag.Int("cpuNum", 0, "num of cpus to use")

	flag.Parse()

	if *name == "" {
		log.Fatal("input a name")
	}

	if *age == 0 {
		log.Fatal("input an age")
	}

	if *cpuNum != 0 {
		runtime.GOMAXPROCS(*cpuNum)
	}

	hello_world_package.OutputName(*name)
	age_package.OutputAge(*age)
}
