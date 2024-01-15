package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DisassociatePublicipsRequest Request Object
type DisassociatePublicipsRequest struct {

	// 弹性公网IP的ID
	PublicipId string `json:"publicip_id"`
}

func (o DisassociatePublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociatePublicipsRequest struct{}"
	}

	return strings.Join([]string{"DisassociatePublicipsRequest", string(data)}, " ")
}
