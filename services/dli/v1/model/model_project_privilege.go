package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ProjectPrivilege struct {
	Object *string `json:"object,omitempty"`

	ApplicantProjectId *string `json:"applicant_project_id,omitempty"`

	Privileges *[]string `json:"privileges,omitempty"`
}

func (o ProjectPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectPrivilege struct{}"
	}

	return strings.Join([]string{"ProjectPrivilege", string(data)}, " ")
}
