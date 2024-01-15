package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBrokersResponse struct {
	Brokers        *[]ListBrokersRespBrokers `json:"brokers,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListBrokersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBrokersResponse struct{}"
	}

	return strings.Join([]string{"ListBrokersResponse", string(data)}, " ")
}
