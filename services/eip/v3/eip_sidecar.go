package v3

import "github.com/fecloud-sdk/fecloud-sdk-go/services/eip/v3/model"

func (c *EipClient) GetListBandwidthRequest() *model.ListBandwidthRequest {
	return new(model.ListBandwidthRequest)
}

func (c *EipClient) GetListCommonPoolsRequest() *model.ListCommonPoolsRequest {
	return new(model.ListCommonPoolsRequest)
}

func (c *EipClient) GetListPublicBorderGroupsRequest() *model.ListPublicBorderGroupsRequest {
	return new(model.ListPublicBorderGroupsRequest)
}

func (c *EipClient) GetListShareBandwidthTypesRequest() *model.ListShareBandwidthTypesRequest {
	return new(model.ListShareBandwidthTypesRequest)
}

func (c *EipClient) GetAssociatePublicipsRequest() *model.AssociatePublicipsRequest {
	return new(model.AssociatePublicipsRequest)
}

func (c *EipClient) GetAttachBatchPublicIpRequest() *model.AttachBatchPublicIpRequest {
	return new(model.AttachBatchPublicIpRequest)
}

func (c *EipClient) GetAttachShareBandwidthRequest() *model.AttachShareBandwidthRequest {
	return new(model.AttachShareBandwidthRequest)
}

func (c *EipClient) GetCountEipAvailableResourcesRequest() *model.CountEipAvailableResourcesRequest {
	return new(model.CountEipAvailableResourcesRequest)
}

func (c *EipClient) GetDetachBatchPublicIpRequest() *model.DetachBatchPublicIpRequest {
	return new(model.DetachBatchPublicIpRequest)
}

func (c *EipClient) GetDetachShareBandwidthRequest() *model.DetachShareBandwidthRequest {
	return new(model.DetachShareBandwidthRequest)
}

func (c *EipClient) GetDisableNat64Request() *model.DisableNat64Request {
	return new(model.DisableNat64Request)
}

func (c *EipClient) GetDisassociatePublicipsRequest() *model.DisassociatePublicipsRequest {
	return new(model.DisassociatePublicipsRequest)
}

func (c *EipClient) GetEnableNat64Request() *model.EnableNat64Request {
	return new(model.EnableNat64Request)
}

func (c *EipClient) GetListPublicipsRequest() *model.ListPublicipsRequest {
	return new(model.ListPublicipsRequest)
}

func (c *EipClient) GetShowPublicipRequest() *model.ShowPublicipRequest {
	return new(model.ShowPublicipRequest)
}
