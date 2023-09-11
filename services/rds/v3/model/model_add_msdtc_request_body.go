package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AddMsdtcRequestBody struct {

	// 主机信息，key为hostname ，value 为IP
	Hosts *[]MsdtcHostOption `json:"hosts,omitempty"`
}

func (o AddMsdtcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMsdtcRequestBody struct{}"
	}

	return strings.Join([]string{"AddMsdtcRequestBody", string(data)}, " ")
}