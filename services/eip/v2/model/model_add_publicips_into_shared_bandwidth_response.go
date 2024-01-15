package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AddPublicipsIntoSharedBandwidthResponse Response Object
type AddPublicipsIntoSharedBandwidthResponse struct {
	Bandwidth      *BandwidthRespInsert `json:"bandwidth,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o AddPublicipsIntoSharedBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPublicipsIntoSharedBandwidthResponse struct{}"
	}

	return strings.Join([]string{"AddPublicipsIntoSharedBandwidthResponse", string(data)}, " ")
}
