package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowInstanceExtendProductInfoResponse struct {
	Hourly *[]ShowInstanceExtendProductInfoRespHourly `json:"hourly,omitempty"`

	Monthly        *[]ShowInstanceExtendProductInfoRespHourly `json:"monthly,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ShowInstanceExtendProductInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceExtendProductInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceExtendProductInfoResponse", string(data)}, " ")
}
