package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteLogGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLogGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogGroupResponse", string(data)}, " ")
}
