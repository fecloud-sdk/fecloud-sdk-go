package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteResourceInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceInstanceTagResponse", string(data)}, " ")
}
