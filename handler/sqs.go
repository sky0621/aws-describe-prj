package handler

import "html/template"

type SqsHandler struct {
}

func (h *SqsHandler) Handle() error {
	template.New("../template/sqs.md")
	return nil
}
