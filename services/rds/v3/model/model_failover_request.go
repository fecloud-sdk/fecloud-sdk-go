package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type FailoverRequest struct {
	Force *bool `json:"force,omitempty"`
}

func (o FailoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailoverRequest struct{}"
	}

	return strings.Join([]string{"FailoverRequest", string(data)}, " ")
}
