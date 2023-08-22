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
	spiderNamePtr := flag.String("n", os.Getenv("CRAWLER_SPIDER_NAME"), "spider name")
	startFuncPtr := flag.String("f", os.Getenv("CRAWLER_SPIDER_FUNC"), "spider func")
	argsPtr := flag.String("a", os.Getenv("CRAWLER_SPIDER_ARGS"), "spider args")
	modePtr := flag.String("m", os.Getenv("CRAWLER_SPIDER_MODE"), "spider mode")
	specPtr := flag.String("s", os.Getenv("CRAWLER_SPIDER_SPEC"), "spider spec")

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
		mode = "manual"
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
