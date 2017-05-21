package structure

type Ec2Information struct {
	Reservations []*Reservation
}

type Reservation struct {
	InstanceID, InstanceType, PublicDnsName, PublicIpAddress, PrivateDnsName, PrivateIpAddress, InstanceState string
}
