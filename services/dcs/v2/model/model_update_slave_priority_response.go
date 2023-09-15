package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateSlavePriorityResponse Response Object
type UpdateSlavePriorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlavePriorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlavePriorityResponse struct{}"
	}

	return strings.Join([]string{"UpdateSlavePriorityResponse", string(data)}, " ")
}
