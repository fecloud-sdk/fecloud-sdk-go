package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLoadbalancerRequest struct {
	Body *CreateLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o CreateLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerRequest", string(data)}, " ")
}
