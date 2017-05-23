RDS

| usecase | environment | instance type | DB name | address | port | engine | version | username | state |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
{{range .}}| {{.Usecase}} | {{.Environment}} | {{.DBInstanceClass}} | {{.DBName}} | {{.EndpointAddress}} | {{.EndpointPort}} | {{.Engine}} | {{.EngineVersion}} | {{.MasterUsername}} | {{.DBInstanceStatus}} |
{{end}}
