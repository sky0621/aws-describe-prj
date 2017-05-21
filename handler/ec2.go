package handler

import (
	"bytes"
	"html/template"

	fs_aws "github.com/sky0621/aws-describe-prj/aws"
	"github.com/sky0621/aws-describe-prj/config"
	"github.com/sky0621/aws-describe-prj/structure"
)

type Ec2Handler struct {
}

func (h *Ec2Handler) Handle() (output *bytes.Buffer, err error) {
	sess, err := fs_aws.NewSession()
	if err != nil {
		return nil, err
	}

	// とりあえずEc2の情報
	info, err := fs_aws.GetEc2Information(fs_aws.NewEc2(sess))
	if err != nil {
		return nil, err
	}

	// 表示に付け足す情報を設定ファイルから取得
	conf := config.NewEc2Config()

	// マージ
	desc := mergeEc2Information(info.Reservations, conf)

	// 表示のためのテンプレートを取得して適用
	tmpl := template.Must(template.ParseFiles(conf.Template))
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, desc)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type Ec2Description struct {
	Type, Environment, InstanceType, PublicDnsName, PublicIpAddress, PrivateDnsName, PrivateIpAddress, InstanceState string
}

// TODO 要リファクタ
func mergeEc2Information(reservations []*structure.Reservation, conf *config.Ec2Config) []Ec2Description {
	descs := []Ec2Description{}
	//f := conf.Filter
	//inFilter := regexp.MustCompile(f.In)
	//var outFilters []*regexp.Regexp
	//for _, out := range f.Out {
	//	outFilters = append(outFilters, regexp.MustCompile(out))
	//}
	for _, res := range reservations {
		//if f.In != "" && inFilter.FindString(qname) == "" {
		//	fmt.Printf("Not match in filter: %v\n", qname)
		//	continue
		//}
		//nodisp := false
		//for _, filter := range outFilters {
		//	if filter.FindString(qname) != "" {
		//		fmt.Printf("Match out filter: %v\n", qname)
		//		nodisp = true
		//		break
		//	}
		//}
		//if nodisp {
		//	continue
		//}
		s := conf.Supplements[res.InstanceID]
		desc := Ec2Description{
			Type:             s.Usecase,
			Environment:      s.Environment,
			InstanceType:     res.InstanceType,
			PublicDnsName:    res.PublicDnsName,
			PublicIpAddress:  res.PublicIpAddress,
			PrivateDnsName:   res.PrivateDnsName,
			PrivateIpAddress: res.PrivateIpAddress,
			InstanceState:    res.InstanceState,
		}
		descs = append(descs, desc)
	}
	return descs
}
