package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteHealthMonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorResponse", string(data)}, " ")
}
