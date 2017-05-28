package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/sky0621/aws-describe-prj/config"
	"github.com/sky0621/aws-describe-prj/subcommand/dynamodb"
	"github.com/sky0621/aws-describe-prj/subcommand/ec2"
	"github.com/sky0621/aws-describe-prj/subcommand/rds"
	"github.com/sky0621/aws-describe-prj/subcommand/sqs"
	"github.com/spiegel-im-spiegel/gofacade"
)

const (
	Name    string = "awsdescribe"
	Version string = "0.1.0"
)

func main() {
	f := flag.String("f", "../../config/config.toml", "Config File Fullpath")
	flag.Parse()

	// Viperグローバル持ち
	err := config.ReadConfig(*f)
	if err != nil {
		panic(err)
	}

	cxt := gofacade.NewContext(os.Stdin, os.Stdout, os.Stderr)
	fcd := setupFacade(cxt)
	rtn, err := fcd.Run(Name, Version, os.Args[1:])
	if err != nil {
		cxt.Error(fmt.Sprintln(err))
	}
	os.Exit(rtn)
}

func setupFacade(cxt *gofacade.Context) *gofacade.Facade {
	fcd := gofacade.NewFacade(cxt)
	fcd.AddCommand(sqs.Name, sqs.Command(cxt, Name))
	fcd.AddCommand(ec2.Name, ec2.Command(cxt, Name))
	fcd.AddCommand(rds.Name, rds.Command(cxt, Name))
	fcd.AddCommand(dynamodb.Name, dynamodb.Command(cxt, Name))
	return fcd
}
