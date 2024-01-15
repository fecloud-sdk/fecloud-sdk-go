package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CancelSqlJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSqlJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSqlJobResponse struct{}"
	}

	return strings.Join([]string{"CancelSqlJobResponse", string(data)}, " ")
}
