package structure

type RdsInformation struct {
	Instances []*RdsInstance
}

type RdsInstance struct {
	DBInstanceClass, DBName, EndpointAddress, Engine, EngineVersion, MasterUsername, DBInstanceStatus string
	EndpointPort                                                                                      int64
}
