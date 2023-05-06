package cli

import (
	"errors"
	"flag"
	"log"
)

type Cli struct {
	ConfigFile string
	StartFunc  string
	Args       string
	Mode       string
}

func NewCli() (c *Cli, err error) {
	configFilePtr := flag.String("c", "", "config file")
	startFuncPtr := flag.String("f", "Test", "start func")
	argsPtr := flag.String("a", "", "args")
	modePtr := flag.String("m", "test", "mode")

	flag.Parse()

	startFunc := *startFuncPtr
	if startFunc == "" {
		err = errors.New("start func is empty")
		return
	}
	log.Printf("func=%s\n", startFunc)

	c = &Cli{
		ConfigFile: *configFilePtr,
		StartFunc:  startFunc,
		Args:       *argsPtr,
		Mode:       *modePtr,
	}

	return
}
