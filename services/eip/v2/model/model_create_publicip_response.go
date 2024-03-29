package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreatePublicipResponse Response Object
type CreatePublicipResponse struct {
	Publicip       *PublicipCreateResp `json:"publicip,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreatePublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicipResponse", string(data)}, " ")
}
