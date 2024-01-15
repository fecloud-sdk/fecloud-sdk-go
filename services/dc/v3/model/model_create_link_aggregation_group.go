package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateLinkAggregationGroup struct {
	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	RegionId *string `json:"region_id,omitempty"`

	WorkMode CreateLinkAggregationGroupWorkMode `json:"work_mode"`

	MinUpNum *int32 `json:"min_up_num,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	DirectConnectIds []string `json:"direct_connect_ids"`

	IesId *string `json:"ies_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateLinkAggregationGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLinkAggregationGroup struct{}"
	}

	return strings.Join([]string{"CreateLinkAggregationGroup", string(data)}, " ")
}

type CreateLinkAggregationGroupWorkMode struct {
	value string
}

type CreateLinkAggregationGroupWorkModeEnum struct {
	MANUAL CreateLinkAggregationGroupWorkMode
	STATIC CreateLinkAggregationGroupWorkMode
}

func GetCreateLinkAggregationGroupWorkModeEnum() CreateLinkAggregationGroupWorkModeEnum {
	return CreateLinkAggregationGroupWorkModeEnum{
		MANUAL: CreateLinkAggregationGroupWorkMode{
			value: "Manual",
		},
		STATIC: CreateLinkAggregationGroupWorkMode{
			value: "Static",
		},
	}
}

func (c CreateLinkAggregationGroupWorkMode) Value() string {
	return c.value
}

func (c CreateLinkAggregationGroupWorkMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLinkAggregationGroupWorkMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
