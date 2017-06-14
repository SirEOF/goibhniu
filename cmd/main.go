package main

import (
	"flag"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/ccdc/gooby/lib/logrus_filehook"
	"github.com/emperorcow/goibhniu/cmd/api"
	"github.com/emperorcow/goibhniu/common"
	"github.com/go-ini/ini"
)

func main() {
	data := common.FermentationData{
		Brews:      common.NewFermentables(),
		Fermenters: common.NewVessels(),
	}

	log.SetOutput(os.Stderr)

	confPath := flag.String("conf", "", "File path to the configuration file.")
	flag.Parse()

	if *confPath == "" {
		log.Error("A configuration file was not included on the command line.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
	}, *confPath)
	if err != nil {
		log.WithField("msg", err.Error()).Fatal("Unable to load configuration file.")
	}

	/////////////////////////////// GENERAL CONFIG //////////////////////////////
	secGeneral, err := cfg.GetSection("General")
	if err != nil {
		log.Fatal("A general configuration section was not configured")
	}

	tmpListenPort := secGeneral.Key("ListenPort").String()
	if tmpListenPort == "" {
		log.Fatal("A listen port was not defined for the API and website.")
	}
	globalListenPort, err := strconv.Atoi(tmpListenPort)
	if err != nil {
		log.Fatal("Unable to successfully convert ListenPort to a number.")
	}

	globalWebRoot := secGeneral.Key("WebRoot").String()
	if globalWebRoot == "" {
		log.Fatal("A web root was not configured for the web files.")
	}

	globalListenAddress := secGeneral.Key("ListenAddress").String()
	if globalListenAddress == "" {
		log.Fatal("The network listen address was not configured under general.")
	}

	switch secGeneral.Key("LogLevel").String() {
	case "Debug":
		log.SetLevel(log.DebugLevel)
	case "Info":
		log.SetLevel(log.InfoLevel)
	case "Warn":
		log.SetLevel(log.WarnLevel)
	case "Error":
		log.SetLevel(log.ErrorLevel)
	case "Fatal":
		log.SetLevel(log.FatalLevel)
	case "Panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	globalLogFile := secGeneral.Key("LogFile").String()
	if globalLogFile != "" {
		hook, err := logrus_filehook.NewFileHook(globalLogFile)
		if err != nil {
			log.WithField("msg", err.Error()).Fatal("Unable to open log file.")
		} else {
			log.AddHook(hook)
		}
	}

	api := api.NewAPIServer(&data, globalListenPort, globalListenAddress)
	api.Serve(globalWebRoot)
}
