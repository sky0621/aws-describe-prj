package structure

type Ec2Information struct {
	Reservations []*Reservation
}

type Reservation struct {
	InstanceType, PublicDnsName, PublicIpAddress, PrivateDnsName, PrivateIpAddress, InstanceState string
}
