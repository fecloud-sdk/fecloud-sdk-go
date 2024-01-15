package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchTestConnectionReq struct {
	Jobs []TestEndPoint `json:"jobs"`
}

func (o BatchTestConnectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTestConnectionReq struct{}"
	}

	return strings.Join([]string{"BatchTestConnectionReq", string(data)}, " ")
}
