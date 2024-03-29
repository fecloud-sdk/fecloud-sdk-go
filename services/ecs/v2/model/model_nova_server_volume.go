package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NovaServerVolume struct {
	Id string `json:"id"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o NovaServerVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServerVolume struct{}"
	}

	return strings.Join([]string{"NovaServerVolume", string(data)}, " ")
}
