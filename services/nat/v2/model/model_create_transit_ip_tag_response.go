package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateTransitIpTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTransitIpTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTransitIpTagResponse", string(data)}, " ")
}
