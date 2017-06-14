package config

import "github.com/spf13/viper"

// Config ...
type Config struct {
	Aws *AwsConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Aws: NewAwsConfig(),
	}
}

type AwsConfig struct {
	Sqs *SqsConfig
}

func NewAwsConfig() *AwsConfig {
	return &AwsConfig{
		Sqs: NewSqsConfig(),
	}
}

type SqsConfig struct {
	Template    string
	Filter      Filter
	Supplements map[string]Supplement
}

func NewSqsConfig() *SqsConfig {
	var s map[string]Supplement
	err := viper.UnmarshalKey("aws.sqs.supplement", &s)
	if err != nil {
		panic(err)
	}
	var f Filter
	err = viper.UnmarshalKey("aws.sqs.filter", &f)
	if err != nil {
		panic(err)
	}
	return &SqsConfig{
		Template:    viper.GetString("aws.sqs.template"),
		Filter:      f,
		Supplements: s,
	}
}

type Ec2Config struct {
	Template    string
	Filter      Filter
	Supplements map[string]Supplement
}

func NewEc2Config() *Ec2Config {
	var s map[string]Supplement
	err := viper.UnmarshalKey("aws.ec2.supplement", &s)
	if err != nil {
		panic(err)
	}
	var f Filter
	err = viper.UnmarshalKey("aws.ec2.filter", &f)
	if err != nil {
		panic(err)
	}
	return &Ec2Config{
		Template:    viper.GetString("aws.ec2.template"),
		Filter:      f,
		Supplements: s,
	}
}

type RdsConfig struct {
	Template    string
	Filter      Filter
	Supplements map[string]Supplement
}

func NewRdsConfig() *RdsConfig {
	var s map[string]Supplement
	err := viper.UnmarshalKey("aws.rds.supplement", &s)
	if err != nil {
		panic(err)
	}
	var f Filter
	err = viper.UnmarshalKey("aws.rds.filter", &f)
	if err != nil {
		panic(err)
	}
	return &RdsConfig{
		Template:    viper.GetString("aws.rds.template"),
		Filter:      f,
		Supplements: s,
	}
}

type DynamoDBConfig struct {
	Template    string
	Filter      Filter
	Supplements map[string]Supplement
}

func NewDynamoDBConfig() *DynamoDBConfig {
	var s map[string]Supplement
	err := viper.UnmarshalKey("aws.dynamodb.supplement", &s)
	if err != nil {
		panic(err)
	}
	var f Filter
	err = viper.UnmarshalKey("aws.dynamodb.filter", &f)
	if err != nil {
		panic(err)
	}
	return &DynamoDBConfig{
		Template:    viper.GetString("aws.dynamodb.template"),
		Filter:      f,
		Supplements: s,
	}
}

type Filter struct {
	In  string
	Out []string
}

type Supplement struct {
	Usecase, Environment string
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
