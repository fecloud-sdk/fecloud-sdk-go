package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CinderShowVolumeTransferRequest struct {
	TransferId string `json:"transfer_id"`
}

func (o CinderShowVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderShowVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderShowVolumeTransferRequest", string(data)}, " ")
}
