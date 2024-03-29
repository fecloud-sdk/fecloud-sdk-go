package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowHealthmonitorsRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthmonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthmonitorsRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthmonitorsRequest", string(data)}, " ")
}
