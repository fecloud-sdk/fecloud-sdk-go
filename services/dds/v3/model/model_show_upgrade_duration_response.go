package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowUpgradeDurationResponse Response Object
type ShowUpgradeDurationResponse struct {

	// 升级策略列表
	Strategies     *[]DurationStrategies `json:"strategies,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowUpgradeDurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeDurationResponse struct{}"
	}

	return strings.Join([]string{"ShowUpgradeDurationResponse", string(data)}, " ")
}
