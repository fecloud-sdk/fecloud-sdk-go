package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLogStreamResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogStreamResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogStreamResponse", string(data)}, " ")
}
