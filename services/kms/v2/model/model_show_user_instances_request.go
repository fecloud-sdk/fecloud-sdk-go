package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type ShowUserInstancesRequest struct {
}

func (o ShowUserInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowUserInstancesRequest", string(data)}, " ")
}
