package handler

import "bytes"

type AwsHandler interface {
	Handle() (output *bytes.Buffer, err error)
}
