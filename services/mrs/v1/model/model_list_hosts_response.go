package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHostsResponse struct {
	Hosts *[]HostModel `json:"hosts,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsResponse struct{}"
	}

	return strings.Join([]string{"ListHostsResponse", string(data)}, " ")
}
