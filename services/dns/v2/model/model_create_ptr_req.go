package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreatePtrReq struct {
	Ptrdname string `json:"ptrdname"`

	Description *string `json:"description,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreatePtrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePtrReq struct{}"
	}

	return strings.Join([]string{"CreatePtrReq", string(data)}, " ")
}
