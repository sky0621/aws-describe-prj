package handler

import (
	"bytes"
	"html/template"

	"strings"

	"regexp"

	fs_aws "github.com/sky0621/aws-describe-prj/aws"
	"github.com/sky0621/aws-describe-prj/config"
)

type DynamoDBHandler struct {
}

func (h *DynamoDBHandler) Handle() (output *bytes.Buffer, err error) {
	sess, err := fs_aws.NewSession()
	if err != nil {
		return nil, err
	}

	// とりあえずDynamoDBの情報
	info, err := fs_aws.GetDynamoDBInformation(fs_aws.NewDynamoDB(sess))
	if err != nil {
		return nil, err
	}

	// 表示に付け足す情報を設定ファイルから取得
	conf := config.NewDynamoDBConfig()

	// マージ
	desc := mergeDynamoDBInformation(info.TableNames, conf)

	// 表示のためのテンプレートを取得して適用
	tmpl := template.Must(template.ParseFiles(conf.Template))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, desc)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type DynamoDBDescription struct {
	Usecase, Environment, TableName string
}

// TODO 要リファクタ
func mergeDynamoDBInformation(tableNames []*string, conf *config.DynamoDBConfig) []DynamoDBDescription {
	descs := []DynamoDBDescription{}
	f := conf.Filter
	inFilter := regexp.MustCompile(f.In)
	var outFilters []*regexp.Regexp
	for _, out := range f.Out {
		outFilters = append(outFilters, regexp.MustCompile(out))
	}
	for _, tableName := range tableNames {
		urlSeps := strings.Split(*tableName, "/")
		qname := urlSeps[len(urlSeps)-1]
		if f.In != "" && inFilter.FindString(qname) == "" {
			//fmt.Printf("Not match in filter: %v\n", qname)
			continue
		}
		nodisp := false
		for _, filter := range outFilters {
			if filter.FindString(qname) != "" {
				//fmt.Printf("Match out filter: %v\n", qname)
				nodisp = true
				break
			}
		}
		if nodisp {
			continue
		}
		s := conf.Supplements[qname]
		desc := DynamoDBDescription{
			Usecase:     s.Usecase,
			Environment: s.Environment,
			TableName:   qname,
		}
		descs = append(descs, desc)
	}
	return descs
}
