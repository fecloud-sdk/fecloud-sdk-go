package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowLoadbalancersStatusResponse struct {
	Statuses       *StatusResp `json:"statuses,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowLoadbalancersStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancersStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancersStatusResponse", string(data)}, " ")
}
