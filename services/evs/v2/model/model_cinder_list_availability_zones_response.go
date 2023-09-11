package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CinderListAvailabilityZonesResponse Response Object
type CinderListAvailabilityZonesResponse struct {

	// 查询请求返回的可用分区列表。
	AvailabilityZoneInfo *[]AzInfo `json:"availabilityZoneInfo,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o CinderListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"CinderListAvailabilityZonesResponse", string(data)}, " ")
}
