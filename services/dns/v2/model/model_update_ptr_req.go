package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdatePtrReq struct {
	Ptrdname string `json:"ptrdname"`

	Description *string `json:"description,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o UpdatePtrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePtrReq struct{}"
	}

	return strings.Join([]string{"UpdatePtrReq", string(data)}, " ")
}
