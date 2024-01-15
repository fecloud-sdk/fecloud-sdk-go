package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListListenersResponse struct {
	Listeners      *[]ListenerResp `json:"listeners,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListListenersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersResponse struct{}"
	}

	return strings.Join([]string{"ListListenersResponse", string(data)}, " ")
}
