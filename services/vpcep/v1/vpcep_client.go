package v1

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/vpcep/v1/model"
)

type VpcepClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVpcepClient(hcClient *http_client.HcHttpClient) *VpcepClient {
	return &VpcepClient{HcClient: hcClient}
}

func VpcepClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *VpcepClient) AcceptOrRejectEndpoint(request *model.AcceptOrRejectEndpointRequest) (*model.AcceptOrRejectEndpointResponse, error) {
	requestDef := GenReqDefForAcceptOrRejectEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptOrRejectEndpointResponse), nil
	}
}

func (c *VpcepClient) AcceptOrRejectEndpointInvoker(request *model.AcceptOrRejectEndpointRequest) *AcceptOrRejectEndpointInvoker {
	requestDef := GenReqDefForAcceptOrRejectEndpoint()
	return &AcceptOrRejectEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) AddOrRemoveServicePermissions(request *model.AddOrRemoveServicePermissionsRequest) (*model.AddOrRemoveServicePermissionsResponse, error) {
	requestDef := GenReqDefForAddOrRemoveServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddOrRemoveServicePermissionsResponse), nil
	}
}

func (c *VpcepClient) AddOrRemoveServicePermissionsInvoker(request *model.AddOrRemoveServicePermissionsRequest) *AddOrRemoveServicePermissionsInvoker {
	requestDef := GenReqDefForAddOrRemoveServicePermissions()
	return &AddOrRemoveServicePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) BatchAddEndpointServicePermissions(request *model.BatchAddEndpointServicePermissionsRequest) (*model.BatchAddEndpointServicePermissionsResponse, error) {
	requestDef := GenReqDefForBatchAddEndpointServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddEndpointServicePermissionsResponse), nil
	}
}

func (c *VpcepClient) BatchAddEndpointServicePermissionsInvoker(request *model.BatchAddEndpointServicePermissionsRequest) *BatchAddEndpointServicePermissionsInvoker {
	requestDef := GenReqDefForBatchAddEndpointServicePermissions()
	return &BatchAddEndpointServicePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) BatchRemoveEndpointServicePermissions(request *model.BatchRemoveEndpointServicePermissionsRequest) (*model.BatchRemoveEndpointServicePermissionsResponse, error) {
	requestDef := GenReqDefForBatchRemoveEndpointServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveEndpointServicePermissionsResponse), nil
	}
}

func (c *VpcepClient) BatchRemoveEndpointServicePermissionsInvoker(request *model.BatchRemoveEndpointServicePermissionsRequest) *BatchRemoveEndpointServicePermissionsInvoker {
	requestDef := GenReqDefForBatchRemoveEndpointServicePermissions()
	return &BatchRemoveEndpointServicePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

func (c *VpcepClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) CreateEndpointService(request *model.CreateEndpointServiceRequest) (*model.CreateEndpointServiceResponse, error) {
	requestDef := GenReqDefForCreateEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointServiceResponse), nil
	}
}

func (c *VpcepClient) CreateEndpointServiceInvoker(request *model.CreateEndpointServiceRequest) *CreateEndpointServiceInvoker {
	requestDef := GenReqDefForCreateEndpointService()
	return &CreateEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

func (c *VpcepClient) DeleteEndpointInvoker(request *model.DeleteEndpointRequest) *DeleteEndpointInvoker {
	requestDef := GenReqDefForDeleteEndpoint()
	return &DeleteEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) DeleteEndpointPolicy(request *model.DeleteEndpointPolicyRequest) (*model.DeleteEndpointPolicyResponse, error) {
	requestDef := GenReqDefForDeleteEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointPolicyResponse), nil
	}
}

func (c *VpcepClient) DeleteEndpointPolicyInvoker(request *model.DeleteEndpointPolicyRequest) *DeleteEndpointPolicyInvoker {
	requestDef := GenReqDefForDeleteEndpointPolicy()
	return &DeleteEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) DeleteEndpointService(request *model.DeleteEndpointServiceRequest) (*model.DeleteEndpointServiceResponse, error) {
	requestDef := GenReqDefForDeleteEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointServiceResponse), nil
	}
}

func (c *VpcepClient) DeleteEndpointServiceInvoker(request *model.DeleteEndpointServiceRequest) *DeleteEndpointServiceInvoker {
	requestDef := GenReqDefForDeleteEndpointService()
	return &DeleteEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListEndpointInfoDetails(request *model.ListEndpointInfoDetailsRequest) (*model.ListEndpointInfoDetailsResponse, error) {
	requestDef := GenReqDefForListEndpointInfoDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointInfoDetailsResponse), nil
	}
}

func (c *VpcepClient) ListEndpointInfoDetailsInvoker(request *model.ListEndpointInfoDetailsRequest) *ListEndpointInfoDetailsInvoker {
	requestDef := GenReqDefForListEndpointInfoDetails()
	return &ListEndpointInfoDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListEndpointService(request *model.ListEndpointServiceRequest) (*model.ListEndpointServiceResponse, error) {
	requestDef := GenReqDefForListEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointServiceResponse), nil
	}
}

func (c *VpcepClient) ListEndpointServiceInvoker(request *model.ListEndpointServiceRequest) *ListEndpointServiceInvoker {
	requestDef := GenReqDefForListEndpointService()
	return &ListEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

func (c *VpcepClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

func (c *VpcepClient) ListQuotaDetailsInvoker(request *model.ListQuotaDetailsRequest) *ListQuotaDetailsInvoker {
	requestDef := GenReqDefForListQuotaDetails()
	return &ListQuotaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListServiceConnections(request *model.ListServiceConnectionsRequest) (*model.ListServiceConnectionsResponse, error) {
	requestDef := GenReqDefForListServiceConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceConnectionsResponse), nil
	}
}

func (c *VpcepClient) ListServiceConnectionsInvoker(request *model.ListServiceConnectionsRequest) *ListServiceConnectionsInvoker {
	requestDef := GenReqDefForListServiceConnections()
	return &ListServiceConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListServiceDescribeDetails(request *model.ListServiceDescribeDetailsRequest) (*model.ListServiceDescribeDetailsResponse, error) {
	requestDef := GenReqDefForListServiceDescribeDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceDescribeDetailsResponse), nil
	}
}

func (c *VpcepClient) ListServiceDescribeDetailsInvoker(request *model.ListServiceDescribeDetailsRequest) *ListServiceDescribeDetailsInvoker {
	requestDef := GenReqDefForListServiceDescribeDetails()
	return &ListServiceDescribeDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListServiceDetails(request *model.ListServiceDetailsRequest) (*model.ListServiceDetailsResponse, error) {
	requestDef := GenReqDefForListServiceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceDetailsResponse), nil
	}
}

func (c *VpcepClient) ListServiceDetailsInvoker(request *model.ListServiceDetailsRequest) *ListServiceDetailsInvoker {
	requestDef := GenReqDefForListServiceDetails()
	return &ListServiceDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListServicePermissionsDetails(request *model.ListServicePermissionsDetailsRequest) (*model.ListServicePermissionsDetailsResponse, error) {
	requestDef := GenReqDefForListServicePermissionsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePermissionsDetailsResponse), nil
	}
}

func (c *VpcepClient) ListServicePermissionsDetailsInvoker(request *model.ListServicePermissionsDetailsRequest) *ListServicePermissionsDetailsInvoker {
	requestDef := GenReqDefForListServicePermissionsDetails()
	return &ListServicePermissionsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListServicePublicDetails(request *model.ListServicePublicDetailsRequest) (*model.ListServicePublicDetailsResponse, error) {
	requestDef := GenReqDefForListServicePublicDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePublicDetailsResponse), nil
	}
}

func (c *VpcepClient) ListServicePublicDetailsInvoker(request *model.ListServicePublicDetailsRequest) *ListServicePublicDetailsInvoker {
	requestDef := GenReqDefForListServicePublicDetails()
	return &ListServicePublicDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListSpecifiedVersionDetails(request *model.ListSpecifiedVersionDetailsRequest) (*model.ListSpecifiedVersionDetailsResponse, error) {
	requestDef := GenReqDefForListSpecifiedVersionDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecifiedVersionDetailsResponse), nil
	}
}

func (c *VpcepClient) ListSpecifiedVersionDetailsInvoker(request *model.ListSpecifiedVersionDetailsRequest) *ListSpecifiedVersionDetailsInvoker {
	requestDef := GenReqDefForListSpecifiedVersionDetails()
	return &ListSpecifiedVersionDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListVersionDetails(request *model.ListVersionDetailsRequest) (*model.ListVersionDetailsResponse, error) {
	requestDef := GenReqDefForListVersionDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionDetailsResponse), nil
	}
}

func (c *VpcepClient) ListVersionDetailsInvoker(request *model.ListVersionDetailsRequest) *ListVersionDetailsInvoker {
	requestDef := GenReqDefForListVersionDetails()
	return &ListVersionDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointConnectionsDesc(request *model.UpdateEndpointConnectionsDescRequest) (*model.UpdateEndpointConnectionsDescResponse, error) {
	requestDef := GenReqDefForUpdateEndpointConnectionsDesc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointConnectionsDescResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointConnectionsDescInvoker(request *model.UpdateEndpointConnectionsDescRequest) *UpdateEndpointConnectionsDescInvoker {
	requestDef := GenReqDefForUpdateEndpointConnectionsDesc()
	return &UpdateEndpointConnectionsDescInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointPolicy(request *model.UpdateEndpointPolicyRequest) (*model.UpdateEndpointPolicyResponse, error) {
	requestDef := GenReqDefForUpdateEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointPolicyResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointPolicyInvoker(request *model.UpdateEndpointPolicyRequest) *UpdateEndpointPolicyInvoker {
	requestDef := GenReqDefForUpdateEndpointPolicy()
	return &UpdateEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointRoutetable(request *model.UpdateEndpointRoutetableRequest) (*model.UpdateEndpointRoutetableResponse, error) {
	requestDef := GenReqDefForUpdateEndpointRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointRoutetableResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointRoutetableInvoker(request *model.UpdateEndpointRoutetableRequest) *UpdateEndpointRoutetableInvoker {
	requestDef := GenReqDefForUpdateEndpointRoutetable()
	return &UpdateEndpointRoutetableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointService(request *model.UpdateEndpointServiceRequest) (*model.UpdateEndpointServiceResponse, error) {
	requestDef := GenReqDefForUpdateEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServiceResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointServiceInvoker(request *model.UpdateEndpointServiceRequest) *UpdateEndpointServiceInvoker {
	requestDef := GenReqDefForUpdateEndpointService()
	return &UpdateEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointServiceName(request *model.UpdateEndpointServiceNameRequest) (*model.UpdateEndpointServiceNameResponse, error) {
	requestDef := GenReqDefForUpdateEndpointServiceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServiceNameResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointServiceNameInvoker(request *model.UpdateEndpointServiceNameRequest) *UpdateEndpointServiceNameInvoker {
	requestDef := GenReqDefForUpdateEndpointServiceName()
	return &UpdateEndpointServiceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointServicePermissionDesc(request *model.UpdateEndpointServicePermissionDescRequest) (*model.UpdateEndpointServicePermissionDescResponse, error) {
	requestDef := GenReqDefForUpdateEndpointServicePermissionDesc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServicePermissionDescResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointServicePermissionDescInvoker(request *model.UpdateEndpointServicePermissionDescRequest) *UpdateEndpointServicePermissionDescInvoker {
	requestDef := GenReqDefForUpdateEndpointServicePermissionDesc()
	return &UpdateEndpointServicePermissionDescInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) UpdateEndpointWhite(request *model.UpdateEndpointWhiteRequest) (*model.UpdateEndpointWhiteResponse, error) {
	requestDef := GenReqDefForUpdateEndpointWhite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointWhiteResponse), nil
	}
}

func (c *VpcepClient) UpdateEndpointWhiteInvoker(request *model.UpdateEndpointWhiteRequest) *UpdateEndpointWhiteInvoker {
	requestDef := GenReqDefForUpdateEndpointWhite()
	return &UpdateEndpointWhiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) BatchAddOrRemoveResourceInstance(request *model.BatchAddOrRemoveResourceInstanceRequest) (*model.BatchAddOrRemoveResourceInstanceResponse, error) {
	requestDef := GenReqDefForBatchAddOrRemoveResourceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddOrRemoveResourceInstanceResponse), nil
	}
}

func (c *VpcepClient) BatchAddOrRemoveResourceInstanceInvoker(request *model.BatchAddOrRemoveResourceInstanceRequest) *BatchAddOrRemoveResourceInstanceInvoker {
	requestDef := GenReqDefForBatchAddOrRemoveResourceInstance()
	return &BatchAddOrRemoveResourceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListQueryProjectResourceTags(request *model.ListQueryProjectResourceTagsRequest) (*model.ListQueryProjectResourceTagsResponse, error) {
	requestDef := GenReqDefForListQueryProjectResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryProjectResourceTagsResponse), nil
	}
}

func (c *VpcepClient) ListQueryProjectResourceTagsInvoker(request *model.ListQueryProjectResourceTagsRequest) *ListQueryProjectResourceTagsInvoker {
	requestDef := GenReqDefForListQueryProjectResourceTags()
	return &ListQueryProjectResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *VpcepClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

func (c *VpcepClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
