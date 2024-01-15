package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ChangingTheDelayThresholdRequestBody struct {
	DelayThresholdInKilobytes int32 `json:"delay_threshold_in_kilobytes"`
}

func (o ChangingTheDelayThresholdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangingTheDelayThresholdRequestBody struct{}"
	}

	return strings.Join([]string{"ChangingTheDelayThresholdRequestBody", string(data)}, " ")
}
