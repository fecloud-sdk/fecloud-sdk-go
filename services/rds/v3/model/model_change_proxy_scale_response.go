package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangeProxyScaleResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeProxyScaleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeProxyScaleResponse struct{}"
	}

	return strings.Join([]string{"ChangeProxyScaleResponse", string(data)}, " ")
}
