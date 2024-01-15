package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHealthmonitorResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateHealthmonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorResponse", string(data)}, " ")
}
