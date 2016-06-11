package main

import (
	"flag"
	"github.com/shinofara/stand/cleaner"
	"github.com/shinofara/stand/compressor"
	"github.com/shinofara/stand/config"
)

type Args struct {
	ConfigPath string
}

func main() {
	var args Args

	flag.StringVar(&args.ConfigPath, "conf", "", "-c path to config yaml")
	// コマンドライン引数を解析
	flag.Parse()

	if args.ConfigPath == "" {
		panic("-c is empty")
	}

	cfgs, _ := config.New(args.ConfigPath)

	err := compressor.Compress(cfgs)

	if err != nil {
		panic(err)
	}

	if err := cleaner.Exec(cfgs); err != nil {
		panic(err)
	}
}
