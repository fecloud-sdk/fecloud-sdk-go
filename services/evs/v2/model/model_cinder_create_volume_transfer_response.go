package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CinderCreateVolumeTransferResponse struct {
	Transfer       *CreateVolumeTransferDetail `json:"transfer,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CinderCreateVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderCreateVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderCreateVolumeTransferResponse", string(data)}, " ")
}
