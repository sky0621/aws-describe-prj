package ecr

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spiegel-im-spiegel/gofacade"
)

// Name はコマンド名を定義する
const Name string = "ecr"

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
	return "Get ECR information"
}

// Help はコマンドのヘルプを返す
func (c Context) Help() string {
	helpText := `
Usage: awsdescribe ecr
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
		return gofacade.ExitFailure
	}
	c.Output("OK ECR!")
	c.Output("OK ECR2!")

	return gofacade.ExitSuccess
}
