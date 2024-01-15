package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListWtpProtectHostResponse struct {
	DataList *[]WtpProtectHostResponseInfo `json:"data_list,omitempty"`

	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWtpProtectHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWtpProtectHostResponse struct{}"
	}

	return strings.Join([]string{"ListWtpProtectHostResponse", string(data)}, " ")
}
