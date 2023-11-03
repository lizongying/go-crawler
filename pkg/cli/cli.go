package cli

import (
	"flag"
	"os"
)

type Cli struct {
	ConfigFile string
	SpiderName string
	StartFunc  string
	Args       string
	Mode       string
	Spec       string
}

func NewCli() (c *Cli, err error) {
	configFilePtr := flag.String("c", os.Getenv("CRAWLER_CONFIG_FILE"), "config file")
	spiderNamePtr := flag.String("n", os.Getenv("CRAWLER_NAME"), "spider name")
	startFuncPtr := flag.String("f", os.Getenv("CRAWLER_FUNC"), "start func")
	argsPtr := flag.String("a", os.Getenv("CRAWLER_ARGS"), "args")
	modePtr := flag.String("m", os.Getenv("CRAWLER_MODE"), "mode")
	specPtr := flag.String("s", os.Getenv("CRAWLER_SPEC"), "spec")

	flag.Parse()

	configFile := *configFilePtr
	spiderName := *spiderNamePtr
	startFunc := *startFuncPtr
	if startFunc == "" {
		startFunc = "Test"
	}
	args := *argsPtr
	mode := *modePtr
	if mode == "" {
		mode = "0"
	}

	c = &Cli{
		ConfigFile: configFile,
		SpiderName: spiderName,
		StartFunc:  startFunc,
		Args:       args,
		Mode:       mode,
		Spec:       *specPtr,
	}

	return
}
