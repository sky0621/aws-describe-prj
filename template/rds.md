RDS

| type | environment | instance type | DB name | address | port | engine | version | username | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .}}| {{.Type}} | {{.Environment}} | {{.DBInstanceClass}} | {{.DBName}} | {{.EndpointAddress}} | {{.EndpointPort}} | {{.Engine}} | {{.EngineVersion}} | {{.MasterUsername}} | {{.DBInstanceStatus}} |
{{end}}
