package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeletePoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoolResponse struct{}"
	}

	return strings.Join([]string{"DeletePoolResponse", string(data)}, " ")
}
