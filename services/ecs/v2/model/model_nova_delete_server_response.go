package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NovaDeleteServerResponse Response Object
type NovaDeleteServerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDeleteServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteServerResponse struct{}"
	}

	return strings.Join([]string{"NovaDeleteServerResponse", string(data)}, " ")
}
