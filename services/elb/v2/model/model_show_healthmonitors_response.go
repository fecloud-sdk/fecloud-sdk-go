package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowHealthmonitorsResponse Response Object
type ShowHealthmonitorsResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowHealthmonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthmonitorsResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthmonitorsResponse", string(data)}, " ")
}
