package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePrivateSnatRequest struct {
	Body *CreatePrivateSnatOptionBody `json:"body,omitempty"`
}

func (o CreatePrivateSnatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatRequest", string(data)}, " ")
}
