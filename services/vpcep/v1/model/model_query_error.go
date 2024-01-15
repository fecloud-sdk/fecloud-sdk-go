package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type QueryError struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o QueryError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryError struct{}"
	}

	return strings.Join([]string{"QueryError", string(data)}, " ")
}
