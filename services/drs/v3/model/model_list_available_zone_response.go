package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListAvailableZoneResponse Response Object
type ListAvailableZoneResponse struct {

	// 可用区信息
	AzInfos        *[]AzInfoResp `json:"az_infos,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAvailableZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZoneResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZoneResponse", string(data)}, " ")
}
