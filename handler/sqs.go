package handler

import (
	"bytes"
	"html/template"

	"strings"

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
	desc := mergeSqsInformation(info.QueueURLs, conf.Supplements)

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

func mergeSqsInformation(queueURLs []*string, supplements map[string]config.Supplement) []SqsDescription {
	descs := []SqsDescription{}
	for _, url := range queueURLs {
		urlSeps := strings.Split(*url, "/")
		qname := urlSeps[len(urlSeps)-1]
		s := supplements[qname]
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
