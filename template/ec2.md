EC2

| type | environment | instance type | [public] hostname | [public] IP | [private] hostname | [private] IP | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .}}| {{.Type}} | {{.Environment}} | {{.InstanceType}} | {{.PublicDnsName}} | {{.PublicIpAddress}} | {{.PrivateDnsName}} | {{.PrivateIpAddress}} | {{.InstanceState}} |
{{end}}
