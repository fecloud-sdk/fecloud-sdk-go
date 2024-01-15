package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ResetMessageOffsetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetMessageOffsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetMessageOffsetResponse struct{}"
	}

	return strings.Join([]string{"ResetMessageOffsetResponse", string(data)}, " ")
}
