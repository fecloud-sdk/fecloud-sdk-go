package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowMaintainWindowsResponse struct {
	MaintainWindows *[]MaintainWindowsEntity `json:"maintain_windows,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ShowMaintainWindowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMaintainWindowsResponse struct{}"
	}

	return strings.Join([]string{"ShowMaintainWindowsResponse", string(data)}, " ")
}
