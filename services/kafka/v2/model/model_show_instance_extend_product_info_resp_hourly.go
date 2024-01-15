package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceExtendProductInfoRespHourly struct {
	Name *string `json:"name,omitempty"`

	Version *string `json:"version,omitempty"`

	Values *[]ShowInstanceExtendProductInfoRespValues `json:"values,omitempty"`
}

func (o ShowInstanceExtendProductInfoRespHourly) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoRespHourly struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoRespHourly", string(data)}, " ")
}
