package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateSharedBandwidthRequest Request Object
type CreateSharedBandwidthRequest struct {
	Body *CreateSharedBandwidhRequestBody `json:"body,omitempty"`
}

func (o CreateSharedBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedBandwidthRequest struct{}"
	}

	return strings.Join([]string{"CreateSharedBandwidthRequest", string(data)}, " ")
}
