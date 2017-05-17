package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/sky0621/aws-describe-prj/structure"
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
			reservations = append(reservations, &structure.Reservation{
				InstanceType:     *instance.InstanceType,
				PublicDnsName:    *instance.PublicDnsName,
				PublicIpAddress:  *instance.PublicIpAddress,
				PrivateDnsName:   *instance.PrivateDnsName,
				PrivateIpAddress: *instance.PrivateIpAddress,
				InstanceState:    *instance.State.Name,
			})
		}
	}

	return &structure.Ec2Information{Reservations: reservations}, nil
}
