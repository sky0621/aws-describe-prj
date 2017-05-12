package sqs

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sky0621/aws-describe-prj/handler"
	"github.com/spiegel-im-spiegel/gofacade"
)

// Name はコマンド名を定義する
const Name string = "sqs"

// Context はコマンドのコンテキストを定義する
type Context struct {
	//Embedded gofacade.Context
	*gofacade.Context
	//AppName にはアプリケーション名を格納する
	AppName string
}

// Command は Context のインスタンスを返す
func Command(cxt *gofacade.Context, appName string) *Context {
	return &Context{Context: cxt, AppName: appName}
}

// Synopsis はコマンドの概要を返す
func (c Context) Synopsis() string {
	return "Get sqs information"
}

// Help はコマンドのヘルプを返す
func (c Context) Help() string {
	helpText := `
Usage: awsdescribe sqs
`
	return fmt.Sprintln(strings.TrimSpace(helpText))
}

// Run はコマンドを実行する
func (c Context) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		c.Error(c.Help())
	}
	// Parse commandline flag
	if err := flags.Parse(args); err != nil {
		return gofacade.ExitCodeError
	}

	h := &handler.SqsHandler{}
	output, err := h.Handle()
	if err != nil {
		panic(err)
	}

	c.Output(output.String())

	return gofacade.ExitCodeOK
}
