package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CinderListVolumeTransfersResponse Response Object
type CinderListVolumeTransfersResponse struct {

	// 云硬盘过户记录列表概要。
	Transfers      *[]VolumeTransferSummary `json:"transfers,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CinderListVolumeTransfersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTransfersResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTransfersResponse", string(data)}, " ")
}
