package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Interval int
	Timeout  int
	Urls     []string
	Listen   string
}

var configFlag = flag.String("c", "devmon.yml", "config file")
var helpFlag = flag.Bool("h", false, "Help")
var cfg Config

func config() int {
	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		return -1
	}

	bt, err := ioutil.ReadFile(*configFlag)
	if err == nil {
		err = yaml.Unmarshal(bt, &cfg)
		if err != nil {
			log.Printf("Error parsing config: %s\n", err)
			return -2
		}
		if cfg.Interval == 0 {
			cfg.Interval = 5
		}
		if cfg.Timeout == 0 {
			cfg.Timeout = 5
		}
		if cfg.Listen == "" {
			cfg.Listen = ":9100"
		}
	} else {
		log.Println("Config not found.")
		flag.PrintDefaults()
		return -3
	}

	return 0
}
