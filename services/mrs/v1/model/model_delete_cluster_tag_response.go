package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteClusterTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClusterTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterTagResponse", string(data)}, " ")
}
