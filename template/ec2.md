EC2 Information

| type | environment | hostname | IP | domain | SSL |
| :--- | :--- | :--- | :--- | :--- | :--- |
{{range .}}| {{.Type}} | {{.Environment}} | {{.Hostname}} | {{.IP}} | {{.Domain}} | {{.SSL}} |
{{end}}
