package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// NeutronUpdateFloatingIpResponse Response Object
type NeutronUpdateFloatingIpResponse struct {
	Floatingip     *PostAndPutFloatingIpResp `json:"floatingip,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o NeutronUpdateFloatingIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFloatingIpResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFloatingIpResponse", string(data)}, " ")
}
