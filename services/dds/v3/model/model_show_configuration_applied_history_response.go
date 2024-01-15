package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationAppliedHistoryResponse struct {
	Histories      *[]ApplyHistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowConfigurationAppliedHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationAppliedHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationAppliedHistoryResponse", string(data)}, " ")
}
