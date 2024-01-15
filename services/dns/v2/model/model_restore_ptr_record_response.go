package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type RestorePtrRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestorePtrRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePtrRecordResponse struct{}"
	}

	return strings.Join([]string{"RestorePtrRecordResponse", string(data)}, " ")
}
