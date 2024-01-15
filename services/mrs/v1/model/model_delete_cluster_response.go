package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteClusterResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterResponse", string(data)}, " ")
}
