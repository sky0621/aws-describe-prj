package main

import (
	"fmt"
	"os"

	"github.com/sky0621/aws-describe-prj/sqs"
	"github.com/spiegel-im-spiegel/gofacade"
)

const (
	Name    string = "awsdescribe"
	Version string = "0.1.0"
)

func main() {
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
	return fcd
}
