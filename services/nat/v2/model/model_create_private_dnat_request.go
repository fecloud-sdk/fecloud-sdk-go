package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateDnatRequest struct {
	Body *CreatePrivateDnatOptionBody `json:"body,omitempty"`
}

func (o CreatePrivateDnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateDnatRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateDnatRequest", string(data)}, " ")
}
