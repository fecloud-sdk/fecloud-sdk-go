package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePublicZoneRequest struct {
	Body *CreatePublicZoneReq `json:"body,omitempty"`
}

func (o CreatePublicZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicZoneRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicZoneRequest", string(data)}, " ")
}
