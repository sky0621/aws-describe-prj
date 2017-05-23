package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/sky0621/aws-describe-prj/structure"
	"github.com/sky0621/aws-describe-prj/util"
)

func NewEc2(sess *session.Session) *ec2.EC2 {
	return ec2.New(sess)
}

func GetEc2Information(cli *ec2.EC2) (*structure.Ec2Information, error) {
	output, err := cli.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		return nil, err
	}

	var reservations []*structure.Reservation
	for _, reservation := range output.Reservations {
		for _, instance := range reservation.Instances {
			if instance == nil {
				continue
			}
			//fmt.Printf("%#v\n", *instance)
			instanceName := ""
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					instanceName = util.ToString(tag.Value)
				}
			}
			reservations = append(reservations, &structure.Reservation{
				InstanceID:       util.ToString(instance.InstanceId),
				InstanceName:     instanceName,
				InstanceType:     util.ToString(instance.InstanceType),
				PublicDnsName:    util.ToString(instance.PublicDnsName),
				PublicIpAddress:  util.ToString(instance.PublicIpAddress),
				PrivateDnsName:   util.ToString(instance.PrivateDnsName),
				PrivateIpAddress: util.ToString(instance.PrivateIpAddress),
				InstanceState:    util.ToString(instance.State.Name),
			})
		}
	}

	return &structure.Ec2Information{Reservations: reservations}, nil
}
