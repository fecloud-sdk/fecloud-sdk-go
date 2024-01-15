package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type EndpointResp struct {
	Id *string `json:"id,omitempty"`

	ServiceType *EndpointRespServiceType `json:"service_type,omitempty"`

	Status *EndpointRespStatus `json:"status,omitempty"`

	ActiveStatus *[]string `json:"active_status,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	MarkerId *int32 `json:"marker_id,omitempty"`

	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	EnableDns *bool `json:"enable_dns,omitempty"`

	DnsNames *[]string `json:"dns_names,omitempty"`

	Ip *string `json:"ip,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Error *[]QueryError `json:"error,omitempty"`

	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Routetables *[]string `json:"routetables,omitempty"`

	Description *string `json:"description,omitempty"`

	PolicyStatement *[]PolicyStatement `json:"policy_statement,omitempty"`

	EndpointPoolId *string `json:"endpoint_pool_id,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o EndpointResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointResp struct{}"
	}

	return strings.Join([]string{"EndpointResp", string(data)}, " ")
}

type EndpointRespServiceType struct {
	value string
}

type EndpointRespServiceTypeEnum struct {
	INTERFACE EndpointRespServiceType
	GATEWAY   EndpointRespServiceType
}

func GetEndpointRespServiceTypeEnum() EndpointRespServiceTypeEnum {
	return EndpointRespServiceTypeEnum{
		INTERFACE: EndpointRespServiceType{
			value: "interface",
		},
		GATEWAY: EndpointRespServiceType{
			value: "gateway",
		},
	}
}

func (c EndpointRespServiceType) Value() string {
	return c.value
}

func (c EndpointRespServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointRespServiceType) UnmarshalJSON(b []byte) error {
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

type EndpointRespStatus struct {
	value string
}

type EndpointRespStatusEnum struct {
	PENDING_ACCEPTANCE EndpointRespStatus
	CREATING           EndpointRespStatus
	ACCEPTED           EndpointRespStatus
	REJECTED           EndpointRespStatus
	FAILED             EndpointRespStatus
	DELETING           EndpointRespStatus
}

func GetEndpointRespStatusEnum() EndpointRespStatusEnum {
	return EndpointRespStatusEnum{
		PENDING_ACCEPTANCE: EndpointRespStatus{
			value: "pendingAcceptance",
		},
		CREATING: EndpointRespStatus{
			value: "creating",
		},
		ACCEPTED: EndpointRespStatus{
			value: "accepted",
		},
		REJECTED: EndpointRespStatus{
			value: "rejected",
		},
		FAILED: EndpointRespStatus{
			value: "failed",
		},
		DELETING: EndpointRespStatus{
			value: "deleting",
		},
	}
}

func (c EndpointRespStatus) Value() string {
	return c.value
}

func (c EndpointRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointRespStatus) UnmarshalJSON(b []byte) error {
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
