package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestorePtrReq struct {
	Ptrdname *interface{} `json:"ptrdname"`
}

func (o RestorePtrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrReq struct{}"
	}

	return strings.Join([]string{"RestorePtrReq", string(data)}, " ")
}
