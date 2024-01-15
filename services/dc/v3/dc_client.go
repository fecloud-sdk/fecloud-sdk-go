package v3

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/dc/v3/model"
)

type DcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDcClient(hcClient *http_client.HcHttpClient) *DcClient {
	return &DcClient{HcClient: hcClient}
}

func DcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *DcClient) CreateHostedDirectConnect(request *model.CreateHostedDirectConnectRequest) (*model.CreateHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForCreateHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHostedDirectConnectResponse), nil
	}
}

func (c *DcClient) CreateHostedDirectConnectInvoker(request *model.CreateHostedDirectConnectRequest) *CreateHostedDirectConnectInvoker {
	requestDef := GenReqDefForCreateHostedDirectConnect()
	return &CreateHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) DeleteDirectConnect(request *model.DeleteDirectConnectRequest) (*model.DeleteDirectConnectResponse, error) {
	requestDef := GenReqDefForDeleteDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDirectConnectResponse), nil
	}
}

func (c *DcClient) DeleteDirectConnectInvoker(request *model.DeleteDirectConnectRequest) *DeleteDirectConnectInvoker {
	requestDef := GenReqDefForDeleteDirectConnect()
	return &DeleteDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) DeleteHostedDirectConnect(request *model.DeleteHostedDirectConnectRequest) (*model.DeleteHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForDeleteHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHostedDirectConnectResponse), nil
	}
}

func (c *DcClient) DeleteHostedDirectConnectInvoker(request *model.DeleteHostedDirectConnectRequest) *DeleteHostedDirectConnectInvoker {
	requestDef := GenReqDefForDeleteHostedDirectConnect()
	return &DeleteHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListDirectConnects(request *model.ListDirectConnectsRequest) (*model.ListDirectConnectsResponse, error) {
	requestDef := GenReqDefForListDirectConnects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectConnectsResponse), nil
	}
}

func (c *DcClient) ListDirectConnectsInvoker(request *model.ListDirectConnectsRequest) *ListDirectConnectsInvoker {
	requestDef := GenReqDefForListDirectConnects()
	return &ListDirectConnectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListHostedDirectConnects(request *model.ListHostedDirectConnectsRequest) (*model.ListHostedDirectConnectsResponse, error) {
	requestDef := GenReqDefForListHostedDirectConnects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostedDirectConnectsResponse), nil
	}
}

func (c *DcClient) ListHostedDirectConnectsInvoker(request *model.ListHostedDirectConnectsRequest) *ListHostedDirectConnectsInvoker {
	requestDef := GenReqDefForListHostedDirectConnects()
	return &ListHostedDirectConnectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowDirectConnect(request *model.ShowDirectConnectRequest) (*model.ShowDirectConnectResponse, error) {
	requestDef := GenReqDefForShowDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDirectConnectResponse), nil
	}
}

func (c *DcClient) ShowDirectConnectInvoker(request *model.ShowDirectConnectRequest) *ShowDirectConnectInvoker {
	requestDef := GenReqDefForShowDirectConnect()
	return &ShowDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowHostedDirectConnect(request *model.ShowHostedDirectConnectRequest) (*model.ShowHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForShowHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHostedDirectConnectResponse), nil
	}
}

func (c *DcClient) ShowHostedDirectConnectInvoker(request *model.ShowHostedDirectConnectRequest) *ShowHostedDirectConnectInvoker {
	requestDef := GenReqDefForShowHostedDirectConnect()
	return &ShowHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) UpdateDirectConnect(request *model.UpdateDirectConnectRequest) (*model.UpdateDirectConnectResponse, error) {
	requestDef := GenReqDefForUpdateDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDirectConnectResponse), nil
	}
}

func (c *DcClient) UpdateDirectConnectInvoker(request *model.UpdateDirectConnectRequest) *UpdateDirectConnectInvoker {
	requestDef := GenReqDefForUpdateDirectConnect()
	return &UpdateDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) UpdateHostedDirectConnect(request *model.UpdateHostedDirectConnectRequest) (*model.UpdateHostedDirectConnectResponse, error) {
	requestDef := GenReqDefForUpdateHostedDirectConnect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostedDirectConnectResponse), nil
	}
}

func (c *DcClient) UpdateHostedDirectConnectInvoker(request *model.UpdateHostedDirectConnectRequest) *UpdateHostedDirectConnectInvoker {
	requestDef := GenReqDefForUpdateHostedDirectConnect()
	return &UpdateHostedDirectConnectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) CreateLinkAggregationGroup(request *model.CreateLinkAggregationGroupRequest) (*model.CreateLinkAggregationGroupResponse, error) {
	requestDef := GenReqDefForCreateLinkAggregationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLinkAggregationGroupResponse), nil
	}
}

func (c *DcClient) CreateLinkAggregationGroupInvoker(request *model.CreateLinkAggregationGroupRequest) *CreateLinkAggregationGroupInvoker {
	requestDef := GenReqDefForCreateLinkAggregationGroup()
	return &CreateLinkAggregationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) DeleteLinkAggregationGroup(request *model.DeleteLinkAggregationGroupRequest) (*model.DeleteLinkAggregationGroupResponse, error) {
	requestDef := GenReqDefForDeleteLinkAggregationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLinkAggregationGroupResponse), nil
	}
}

func (c *DcClient) DeleteLinkAggregationGroupInvoker(request *model.DeleteLinkAggregationGroupRequest) *DeleteLinkAggregationGroupInvoker {
	requestDef := GenReqDefForDeleteLinkAggregationGroup()
	return &DeleteLinkAggregationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListLinkAggregationGroup(request *model.ListLinkAggregationGroupRequest) (*model.ListLinkAggregationGroupResponse, error) {
	requestDef := GenReqDefForListLinkAggregationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLinkAggregationGroupResponse), nil
	}
}

func (c *DcClient) ListLinkAggregationGroupInvoker(request *model.ListLinkAggregationGroupRequest) *ListLinkAggregationGroupInvoker {
	requestDef := GenReqDefForListLinkAggregationGroup()
	return &ListLinkAggregationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowLinkAggregationGroup(request *model.ShowLinkAggregationGroupRequest) (*model.ShowLinkAggregationGroupResponse, error) {
	requestDef := GenReqDefForShowLinkAggregationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLinkAggregationGroupResponse), nil
	}
}

func (c *DcClient) ShowLinkAggregationGroupInvoker(request *model.ShowLinkAggregationGroupRequest) *ShowLinkAggregationGroupInvoker {
	requestDef := GenReqDefForShowLinkAggregationGroup()
	return &ShowLinkAggregationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) UpdateLinkAggregationGroup(request *model.UpdateLinkAggregationGroupRequest) (*model.UpdateLinkAggregationGroupResponse, error) {
	requestDef := GenReqDefForUpdateLinkAggregationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLinkAggregationGroupResponse), nil
	}
}

func (c *DcClient) UpdateLinkAggregationGroupInvoker(request *model.UpdateLinkAggregationGroupRequest) *UpdateLinkAggregationGroupInvoker {
	requestDef := GenReqDefForUpdateLinkAggregationGroup()
	return &UpdateLinkAggregationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

func (c *DcClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) BatchCreateResourceTags(request *model.BatchCreateResourceTagsRequest) (*model.BatchCreateResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourceTagsResponse), nil
	}
}

func (c *DcClient) BatchCreateResourceTagsInvoker(request *model.BatchCreateResourceTagsRequest) *BatchCreateResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateResourceTags()
	return &BatchCreateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) CreateResourceTag(request *model.CreateResourceTagRequest) (*model.CreateResourceTagResponse, error) {
	requestDef := GenReqDefForCreateResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagResponse), nil
	}
}

func (c *DcClient) CreateResourceTagInvoker(request *model.CreateResourceTagRequest) *CreateResourceTagInvoker {
	requestDef := GenReqDefForCreateResourceTag()
	return &CreateResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

func (c *DcClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

func (c *DcClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListTagResourceInstances(request *model.ListTagResourceInstancesRequest) (*model.ListTagResourceInstancesResponse, error) {
	requestDef := GenReqDefForListTagResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResourceInstancesResponse), nil
	}
}

func (c *DcClient) ListTagResourceInstancesInvoker(request *model.ListTagResourceInstancesRequest) *ListTagResourceInstancesInvoker {
	requestDef := GenReqDefForListTagResourceInstances()
	return &ListTagResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

func (c *DcClient) ShowResourceTagInvoker(request *model.ShowResourceTagRequest) *ShowResourceTagInvoker {
	requestDef := GenReqDefForShowResourceTag()
	return &ShowResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) CreateVirtualGateway(request *model.CreateVirtualGatewayRequest) (*model.CreateVirtualGatewayResponse, error) {
	requestDef := GenReqDefForCreateVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVirtualGatewayResponse), nil
	}
}

func (c *DcClient) CreateVirtualGatewayInvoker(request *model.CreateVirtualGatewayRequest) *CreateVirtualGatewayInvoker {
	requestDef := GenReqDefForCreateVirtualGateway()
	return &CreateVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) DeleteVirtualGateway(request *model.DeleteVirtualGatewayRequest) (*model.DeleteVirtualGatewayResponse, error) {
	requestDef := GenReqDefForDeleteVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVirtualGatewayResponse), nil
	}
}

func (c *DcClient) DeleteVirtualGatewayInvoker(request *model.DeleteVirtualGatewayRequest) *DeleteVirtualGatewayInvoker {
	requestDef := GenReqDefForDeleteVirtualGateway()
	return &DeleteVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListVirtualGateways(request *model.ListVirtualGatewaysRequest) (*model.ListVirtualGatewaysResponse, error) {
	requestDef := GenReqDefForListVirtualGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVirtualGatewaysResponse), nil
	}
}

func (c *DcClient) ListVirtualGatewaysInvoker(request *model.ListVirtualGatewaysRequest) *ListVirtualGatewaysInvoker {
	requestDef := GenReqDefForListVirtualGateways()
	return &ListVirtualGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowVirtualGateway(request *model.ShowVirtualGatewayRequest) (*model.ShowVirtualGatewayResponse, error) {
	requestDef := GenReqDefForShowVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVirtualGatewayResponse), nil
	}
}

func (c *DcClient) ShowVirtualGatewayInvoker(request *model.ShowVirtualGatewayRequest) *ShowVirtualGatewayInvoker {
	requestDef := GenReqDefForShowVirtualGateway()
	return &ShowVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) UpdateVirtualGateway(request *model.UpdateVirtualGatewayRequest) (*model.UpdateVirtualGatewayResponse, error) {
	requestDef := GenReqDefForUpdateVirtualGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVirtualGatewayResponse), nil
	}
}

func (c *DcClient) UpdateVirtualGatewayInvoker(request *model.UpdateVirtualGatewayRequest) *UpdateVirtualGatewayInvoker {
	requestDef := GenReqDefForUpdateVirtualGateway()
	return &UpdateVirtualGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) CreateVirtualInterface(request *model.CreateVirtualInterfaceRequest) (*model.CreateVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForCreateVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVirtualInterfaceResponse), nil
	}
}

func (c *DcClient) CreateVirtualInterfaceInvoker(request *model.CreateVirtualInterfaceRequest) *CreateVirtualInterfaceInvoker {
	requestDef := GenReqDefForCreateVirtualInterface()
	return &CreateVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) DeleteVirtualInterface(request *model.DeleteVirtualInterfaceRequest) (*model.DeleteVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForDeleteVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVirtualInterfaceResponse), nil
	}
}

func (c *DcClient) DeleteVirtualInterfaceInvoker(request *model.DeleteVirtualInterfaceRequest) *DeleteVirtualInterfaceInvoker {
	requestDef := GenReqDefForDeleteVirtualInterface()
	return &DeleteVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ListVirtualInterfaces(request *model.ListVirtualInterfacesRequest) (*model.ListVirtualInterfacesResponse, error) {
	requestDef := GenReqDefForListVirtualInterfaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVirtualInterfacesResponse), nil
	}
}

func (c *DcClient) ListVirtualInterfacesInvoker(request *model.ListVirtualInterfacesRequest) *ListVirtualInterfacesInvoker {
	requestDef := GenReqDefForListVirtualInterfaces()
	return &ListVirtualInterfacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) ShowVirtualInterface(request *model.ShowVirtualInterfaceRequest) (*model.ShowVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForShowVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVirtualInterfaceResponse), nil
	}
}

func (c *DcClient) ShowVirtualInterfaceInvoker(request *model.ShowVirtualInterfaceRequest) *ShowVirtualInterfaceInvoker {
	requestDef := GenReqDefForShowVirtualInterface()
	return &ShowVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DcClient) UpdateVirtualInterface(request *model.UpdateVirtualInterfaceRequest) (*model.UpdateVirtualInterfaceResponse, error) {
	requestDef := GenReqDefForUpdateVirtualInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVirtualInterfaceResponse), nil
	}
}

func (c *DcClient) UpdateVirtualInterfaceInvoker(request *model.UpdateVirtualInterfaceRequest) *UpdateVirtualInterfaceInvoker {
	requestDef := GenReqDefForUpdateVirtualInterface()
	return &UpdateVirtualInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
