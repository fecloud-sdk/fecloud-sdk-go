package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CinderCreateVolumeTransferRequestBody struct {
	Transfer *CreateVolumeTransferOption `json:"transfer"`
}

func (o CinderCreateVolumeTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderCreateVolumeTransferRequestBody struct{}"
	}

	return strings.Join([]string{"CinderCreateVolumeTransferRequestBody", string(data)}, " ")
}
