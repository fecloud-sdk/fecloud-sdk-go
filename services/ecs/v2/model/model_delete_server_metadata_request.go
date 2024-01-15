package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteServerMetadataRequest struct {
	Key string `json:"key"`

	ServerId string `json:"server_id"`
}

func (o DeleteServerMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerMetadataRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerMetadataRequest", string(data)}, " ")
}
