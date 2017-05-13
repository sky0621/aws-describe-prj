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
	Supplements map[string]Supplement
}

func NewSqsConfig() *SqsConfig {
	var s map[string]Supplement
	err := viper.UnmarshalKey("aws.sqs.supplement", &s)
	if err != nil {
		panic(err)
	}
	return &SqsConfig{
		Template:    viper.GetString("aws.sqs.template"),
		Supplements: s,
	}
}

type Supplement struct {
	Usecase, Environment string
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
