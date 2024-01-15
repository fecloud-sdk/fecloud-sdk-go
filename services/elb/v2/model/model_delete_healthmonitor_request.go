package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHealthmonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorRequest", string(data)}, " ")
}
