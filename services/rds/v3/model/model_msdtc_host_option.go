package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type MsdtcHostOption struct {
	HostName string `json:"host_name"`

	Ip string `json:"ip"`
}

func (o MsdtcHostOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MsdtcHostOption struct{}"
	}

	return strings.Join([]string{"MsdtcHostOption", string(data)}, " ")
}
