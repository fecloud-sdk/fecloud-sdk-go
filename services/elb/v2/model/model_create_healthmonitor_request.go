package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateHealthmonitorRequest struct {
	Body *CreateHealthmonitorRequestBody `json:"body,omitempty"`
}

func (o CreateHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorRequest", string(data)}, " ")
}
