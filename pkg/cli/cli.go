package cli

import (
	"flag"
	"log"
	"os"
)

type Cli struct {
	ConfigFile string
	StartFunc  string
	Args       string
	Mode       string
}

func NewCli() (c *Cli, err error) {
	configFilePtr := flag.String("c", os.Getenv("CRAWLER_CONFIG_FILE"), "config file")
	startFuncPtr := flag.String("f", os.Getenv("CRAWLER_START_FUNC"), "start func")
	argsPtr := flag.String("a", os.Getenv("CRAWLER_ARGS"), "args")
	modePtr := flag.String("m", os.Getenv("CRAWLER_MODE"), "mode")

	flag.Parse()

	configFile := *configFilePtr
	startFunc := *startFuncPtr
	if startFunc == "" {
		startFunc = "Test"
	}
	args := *argsPtr
	mode := *modePtr
	if mode == "" {
		mode = "test"
	}

	log.Printf("func=%s, args=%s\n", startFunc, args)

	c = &Cli{
		ConfigFile: configFile,
		StartFunc:  startFunc,
		Args:       args,
		Mode:       mode,
	}

	return
}
