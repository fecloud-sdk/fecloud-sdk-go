package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMaintainWindowsRequest struct {
}

func (o ShowMaintainWindowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMaintainWindowsRequest struct{}"
	}

	return strings.Join([]string{"ShowMaintainWindowsRequest", string(data)}, " ")
}
