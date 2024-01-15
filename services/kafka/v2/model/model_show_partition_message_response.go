package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowPartitionMessageResponse struct {
	Message        *[]ShowPartitionMessageEntity `json:"message,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowPartitionMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionMessageResponse struct{}"
	}

	return strings.Join([]string{"ShowPartitionMessageResponse", string(data)}, " ")
}
