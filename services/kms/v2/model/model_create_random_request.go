package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateRandomRequest struct {
	Body *GenRandomRequestBody `json:"body,omitempty"`
}

func (o CreateRandomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRandomRequest struct{}"
	}

	return strings.Join([]string{"CreateRandomRequest", string(data)}, " ")
}
