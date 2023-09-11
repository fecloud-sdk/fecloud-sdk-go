package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CinderShowVolumeTransferResponse Response Object
type CinderShowVolumeTransferResponse struct {
	Transfer       *VolumeTransfer `json:"transfer,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CinderShowVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderShowVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderShowVolumeTransferResponse", string(data)}, " ")
}
