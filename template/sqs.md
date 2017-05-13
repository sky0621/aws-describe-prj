SQS Information

| usecase | environment | queue_name | url |
| :--- | :--- | :--- | :--- |
{{range .}}| {{.Usecase}} | {{.Environment}} | {{.QueueName}} | {{.QueueURL}} |
{{end}}
