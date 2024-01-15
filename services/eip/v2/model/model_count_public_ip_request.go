package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CountPublicIpRequest Request Object
type CountPublicIpRequest struct {
}

func (o CountPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPublicIpRequest struct{}"
	}

	return strings.Join([]string{"CountPublicIpRequest", string(data)}, " ")
}
