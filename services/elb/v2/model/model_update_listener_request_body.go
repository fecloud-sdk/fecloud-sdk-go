package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateListenerRequestBody This is a auto create Body Object
type UpdateListenerRequestBody struct {
	Listener *UpdateListenerReq `json:"listener"`
}

func (o UpdateListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateListenerRequestBody", string(data)}, " ")
}