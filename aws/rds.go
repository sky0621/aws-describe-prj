package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/sky0621/aws-describe-prj/structure"
	"github.com/sky0621/aws-describe-prj/util"
)

func NewRds(sess *session.Session) *rds.RDS {
	return rds.New(sess)
}

func GetRdsInformation(cli *rds.RDS) (*structure.RdsInformation, error) {
	output, err := cli.DescribeDBInstances(&rds.DescribeDBInstancesInput{})
	if err != nil {
		return nil, err
	}

	var instances []*structure.RdsInstance
	for _, instance := range output.DBInstances {
		//fmt.Printf("%#v\n", instance)
		instances = append(instances, &structure.RdsInstance{
			DBInstanceClass:  util.ToString(instance.DBInstanceClass),
			DBName:           util.ToString(instance.DBName),
			EndpointAddress:  util.ToString(instance.Endpoint.Address),
			EndpointPort:     util.ToInt64(instance.Endpoint.Port),
			Engine:           util.ToString(instance.Engine),
			EngineVersion:    util.ToString(instance.EngineVersion),
			MasterUsername:   util.ToString(instance.MasterUsername),
			DBInstanceStatus: util.ToString(instance.DBInstanceStatus),
		})
	}

	return &structure.RdsInformation{Instances: instances}, nil
}
