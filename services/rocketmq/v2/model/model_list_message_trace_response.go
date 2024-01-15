package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListMessageTraceResponse struct {
	Trace          *[]ListMessageTraceRespTrace `json:"trace,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMessageTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTraceResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTraceResponse", string(data)}, " ")
}
