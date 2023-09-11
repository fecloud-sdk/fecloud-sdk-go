package v3

import "github.com/fecloud-sdk/fecloud-sdk-go/services/vpc/v3/model"

func (c *VpcClient) GetAddVpcExtendCidrRequest() *model.AddVpcExtendCidrRequest {
	return new(model.AddVpcExtendCidrRequest)
}

func (c *VpcClient) GetCreateVpcRequest() *model.CreateVpcRequest {
	return new(model.CreateVpcRequest)
}

func (c *VpcClient) GetDeleteVpcRequest() *model.DeleteVpcRequest {
	return new(model.DeleteVpcRequest)
}

func (c *VpcClient) GetListVpcsRequest() *model.ListVpcsRequest {
	return new(model.ListVpcsRequest)
}

func (c *VpcClient) GetRemoveVpcExtendCidrRequest() *model.RemoveVpcExtendCidrRequest {
	return new(model.RemoveVpcExtendCidrRequest)
}

func (c *VpcClient) GetShowVpcRequest() *model.ShowVpcRequest {
	return new(model.ShowVpcRequest)
}

func (c *VpcClient) GetUpdateVpcRequest() *model.UpdateVpcRequest {
	return new(model.UpdateVpcRequest)
}
