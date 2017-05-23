EC2

| usecase | environment | instance name | instance type | [public] hostname | [public] IP | [private] hostname | [private] IP | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .}}| {{.Usecase}} | {{.Environment}} | {{.InstanceName}} | {{.InstanceType}} | {{.PublicDnsName}} | {{.PublicIpAddress}} | {{.PrivateDnsName}} | {{.PrivateIpAddress}} | {{.InstanceState}} |
{{end}}
