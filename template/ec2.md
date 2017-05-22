EC2

| usecase | environment | instance type | [public] hostname | [public] IP | [private] hostname | [private] IP | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .}}| {{.Usecase}} | {{.Environment}} | {{.InstanceType}} | {{.PublicDnsName}} | {{.PublicIpAddress}} | {{.PrivateDnsName}} | {{.PrivateIpAddress}} | {{.InstanceState}} |
{{end}}
