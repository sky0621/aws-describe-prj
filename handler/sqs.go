package handler

import (
	"bytes"
	"html/template"

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

	// 表示のためのテンプレートを取得
	tmpl := template.Must(template.ParseFiles(conf.Template))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, info)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
