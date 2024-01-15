package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteSharedBandwidthResponse Response Object
type DeleteSharedBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSharedBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteSharedBandwidthResponse", string(data)}, " ")
}
