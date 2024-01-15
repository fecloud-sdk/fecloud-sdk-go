package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigurationModifyHistoryResponse struct {
	Histories      *[]HistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowConfigurationModifyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigurationModifyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigurationModifyHistoryResponse", string(data)}, " ")
}
