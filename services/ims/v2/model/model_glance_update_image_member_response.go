package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type GlanceUpdateImageMemberResponse struct {
	Status *string `json:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ImageId *string `json:"image_id,omitempty"`

	MemberId *string `json:"member_id,omitempty"`

	Schema         *string `json:"schema,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceUpdateImageMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberResponse struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberResponse", string(data)}, " ")
}
