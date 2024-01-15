package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImageTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageTagResponse", string(data)}, " ")
}
