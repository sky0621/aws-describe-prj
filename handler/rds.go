package handler

import (
	"bytes"
	"html/template"

	"regexp"

	fs_aws "github.com/sky0621/aws-describe-prj/aws"
	"github.com/sky0621/aws-describe-prj/config"
	"github.com/sky0621/aws-describe-prj/structure"
)

type RdsHandler struct {
}

func (h *RdsHandler) Handle() (output *bytes.Buffer, err error) {
	// とりあえずRdsの情報
	info, err := fs_aws.GetRdsInformation(fs_aws.NewRds(fs_aws.NewSession()))
	if err != nil {
		return nil, err
	}

	// 表示に付け足す情報を設定ファイルから取得
	conf := config.NewRdsConfig()

	// マージ
	desc := mergeRdsInformation(info.Instances, conf)
	//fmt.Printf("%#v\n", desc)

	// 表示のためのテンプレートを取得して適用
	tmpl := template.Must(template.ParseFiles(conf.Template))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, desc)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type RdsDescription struct {
	Usecase, Environment, DBInstanceClass, DBName, EndpointAddress, Engine, EngineVersion, MasterUsername, DBInstanceStatus string
	EndpointPort                                                                                                            int64
}

// TODO 要リファクタ
func mergeRdsInformation(instances []*structure.RdsInstance, conf *config.RdsConfig) []RdsDescription {
	descs := []RdsDescription{}
	f := conf.Filter
	inFilter := regexp.MustCompile(f.In)
	var outFilters []*regexp.Regexp
	for _, out := range f.Out {
		outFilters = append(outFilters, regexp.MustCompile(out))
	}
	for _, res := range instances {
		//fmt.Printf("%#v\n", res)
		if f.In != "" && inFilter.FindString(res.EndpointAddress) == "" {
			//fmt.Printf("Not match in filter: %v\n", res.EndpointAddress)
			continue
		}
		nodisp := false
		for _, filter := range outFilters {
			if filter.FindString(res.EndpointAddress) != "" {
				//fmt.Printf("Match out filter: %v\n", res.EndpointAddress)
				nodisp = true
				break
			}
		}
		if nodisp {
			continue
		}
		s := conf.Supplements[res.EndpointAddress]
		desc := RdsDescription{
			Usecase:          s.Usecase,
			Environment:      s.Environment,
			DBInstanceClass:  res.DBInstanceClass,
			DBName:           res.DBName,
			EndpointAddress:  res.EndpointAddress,
			EndpointPort:     res.EndpointPort,
			Engine:           res.Engine,
			EngineVersion:    res.EngineVersion,
			MasterUsername:   res.MasterUsername,
			DBInstanceStatus: res.DBInstanceStatus,
		}
		descs = append(descs, desc)
	}
	return descs
}
