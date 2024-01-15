package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteFlinkJobResponse struct {
	IsSuccess *bool `json:"is_success,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlinkJobResponse", string(data)}, " ")
}
