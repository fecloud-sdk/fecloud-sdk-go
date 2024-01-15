package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"errors"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/converter"

	"strings"
)

type CreateEndpointResponse struct {
	Id *string `json:"id,omitempty"`

	ServiceType *CreateEndpointResponseServiceType `json:"service_type,omitempty"`

	Status *CreateEndpointResponseStatus `json:"status,omitempty"`

	Ip *string `json:"ip,omitempty"`

	ActiveStatus *[]string `json:"active_status,omitempty"`

	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	MarkerId *int32 `json:"marker_id,omitempty"`

	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	EnableDns *bool `json:"enable_dns,omitempty"`

	DnsNames *[]string `json:"dns_names,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Tags *[]TagList `json:"tags,omitempty"`

	Whitelist *[]string `json:"whitelist,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Routetables *[]string `json:"routetables,omitempty"`

	SpecificationName *string `json:"specification_name,omitempty"`

	Description *string `json:"description,omitempty"`

	PolicyStatement *[]string `json:"policy_statement,omitempty"`

	EnableStatus *string `json:"enable_status,omitempty"`

	EndpointPoolId *string `json:"endpoint_pool_id,omitempty"`

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o CreateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointResponse struct{}"
	}

	return strings.Join([]string{"CreateEndpointResponse", string(data)}, " ")
}

type CreateEndpointResponseServiceType struct {
	value string
}

type CreateEndpointResponseServiceTypeEnum struct {
	INTERFACE CreateEndpointResponseServiceType
	GATEWAY   CreateEndpointResponseServiceType
}

func GetCreateEndpointResponseServiceTypeEnum() CreateEndpointResponseServiceTypeEnum {
	return CreateEndpointResponseServiceTypeEnum{
		INTERFACE: CreateEndpointResponseServiceType{
			value: "interface",
		},
		GATEWAY: CreateEndpointResponseServiceType{
			value: "gateway",
		},
	}
}

func (c CreateEndpointResponseServiceType) Value() string {
	return c.value
}

func (c CreateEndpointResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointResponseServiceType) UnmarshalJSON(b []byte) error {
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

type CreateEndpointResponseStatus struct {
	value string
}

type CreateEndpointResponseStatusEnum struct {
	PENDING_ACCEPTANCE CreateEndpointResponseStatus
	CREATING           CreateEndpointResponseStatus
	ACCEPTED           CreateEndpointResponseStatus
	REJECTED           CreateEndpointResponseStatus
	FAILED             CreateEndpointResponseStatus
	DELETING           CreateEndpointResponseStatus
}

func GetCreateEndpointResponseStatusEnum() CreateEndpointResponseStatusEnum {
	return CreateEndpointResponseStatusEnum{
		PENDING_ACCEPTANCE: CreateEndpointResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: CreateEndpointResponseStatus{
			value: "creating",
		},
		ACCEPTED: CreateEndpointResponseStatus{
			value: "accepted",
		},
		REJECTED: CreateEndpointResponseStatus{
			value: "rejected",
		},
		FAILED: CreateEndpointResponseStatus{
			value: "failed",
		},
		DELETING: CreateEndpointResponseStatus{
			value: "deleting",
		},
	}
}

func (c CreateEndpointResponseStatus) Value() string {
	return c.value
}

func (c CreateEndpointResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEndpointResponseStatus) UnmarshalJSON(b []byte) error {
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
