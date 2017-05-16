package handler

import (
	"bytes"
	"html/template"

	"strings"

	"regexp"

	"fmt"

	fs_aws "github.com/sky0621/aws-describe-prj/aws"
	"github.com/sky0621/aws-describe-prj/config"
)

type SqsHandler struct {
}

func (h *SqsHandler) Handle() (output *bytes.Buffer, err error) {
	sess, err := fs_aws.NewSession()
	if err != nil {
		return nil, err
	}

	// とりあえずSQSの情報
	info, err := fs_aws.GetSqsInformation(fs_aws.NewSqs(sess))
	if err != nil {
		return nil, err
	}

	// 表示に付け足す情報を設定ファイルから取得
	conf := config.NewSqsConfig()

	// マージ
	desc := mergeSqsInformation(info.QueueURLs, conf)

	// 表示のためのテンプレートを取得して適用
	tmpl := template.Must(template.ParseFiles(conf.Template))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, desc)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type SqsDescription struct {
	Usecase, Environment, QueueName, QueueURL string
}

// TODO 要リファクタ
func mergeSqsInformation(queueURLs []*string, conf *config.SqsConfig) []SqsDescription {
	descs := []SqsDescription{}
	f := conf.Filter
	inFilter := regexp.MustCompile(f.In)
	var outFilters []*regexp.Regexp
	for _, out := range f.Out {
		outFilters = append(outFilters, regexp.MustCompile(out))
	}
	for _, url := range queueURLs {
		urlSeps := strings.Split(*url, "/")
		qname := urlSeps[len(urlSeps)-1]
		if f.In != "" && inFilter.FindString(qname) == "" {
			fmt.Printf("Not match in filter: %v\n", qname)
			continue
		}
		nodisp := false
		for _, filter := range outFilters {
			if filter.FindString(qname) != "" {
				fmt.Printf("Match out filter: %v\n", qname)
				nodisp = true
				break
			}
		}
		if nodisp {
			continue
		}
		s := conf.Supplements[qname]
		desc := SqsDescription{
			Usecase:     s.Usecase,
			Environment: s.Environment,
			QueueName:   qname,
			QueueURL:    *url,
		}
		descs = append(descs, desc)
	}
	return descs
}
