package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CinderShowVolumeTransferRequest Request Object
type CinderShowVolumeTransferRequest struct {

	// 云硬盘过户记录ID
	TransferId string `json:"transfer_id"`
}

func (o CinderShowVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderShowVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderShowVolumeTransferRequest", string(data)}, " ")
}
