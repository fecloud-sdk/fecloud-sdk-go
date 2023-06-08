package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// Request Object
type ShowVersionsRequest struct {
}

func (o ShowVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionsRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionsRequest", string(data)}, " ")
}
