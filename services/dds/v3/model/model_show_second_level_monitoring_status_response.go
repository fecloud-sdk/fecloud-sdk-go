package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowSecondLevelMonitoringStatusResponse struct {
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSecondLevelMonitoringStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecondLevelMonitoringStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSecondLevelMonitoringStatusResponse", string(data)}, " ")
}
