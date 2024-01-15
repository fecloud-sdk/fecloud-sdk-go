package v2

import (
	http_client "github.com/fecloud-sdk/fecloud-sdk-go/core"
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/dns/v2/model"
)

type DnsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDnsClient(hcClient *http_client.HcHttpClient) *DnsClient {
	return &DnsClient{HcClient: hcClient}
}

func DnsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *DnsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

func (c *DnsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListNameServers(request *model.ListNameServersRequest) (*model.ListNameServersResponse, error) {
	requestDef := GenReqDefForListNameServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNameServersResponse), nil
	}
}

func (c *DnsClient) ListNameServersInvoker(request *model.ListNameServersRequest) *ListNameServersInvoker {
	requestDef := GenReqDefForListNameServers()
	return &ListNameServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowApiInfo(request *model.ShowApiInfoRequest) (*model.ShowApiInfoResponse, error) {
	requestDef := GenReqDefForShowApiInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiInfoResponse), nil
	}
}

func (c *DnsClient) ShowApiInfoInvoker(request *model.ShowApiInfoRequest) *ShowApiInfoInvoker {
	requestDef := GenReqDefForShowApiInfo()
	return &ShowApiInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowDomainQuota(request *model.ShowDomainQuotaRequest) (*model.ShowDomainQuotaResponse, error) {
	requestDef := GenReqDefForShowDomainQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainQuotaResponse), nil
	}
}

func (c *DnsClient) ShowDomainQuotaInvoker(request *model.ShowDomainQuotaRequest) *ShowDomainQuotaInvoker {
	requestDef := GenReqDefForShowDomainQuota()
	return &ShowDomainQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) CreateEipRecordSet(request *model.CreateEipRecordSetRequest) (*model.CreateEipRecordSetResponse, error) {
	requestDef := GenReqDefForCreateEipRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEipRecordSetResponse), nil
	}
}

func (c *DnsClient) CreateEipRecordSetInvoker(request *model.CreateEipRecordSetRequest) *CreateEipRecordSetInvoker {
	requestDef := GenReqDefForCreateEipRecordSet()
	return &CreateEipRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListPtrRecords(request *model.ListPtrRecordsRequest) (*model.ListPtrRecordsResponse, error) {
	requestDef := GenReqDefForListPtrRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPtrRecordsResponse), nil
	}
}

func (c *DnsClient) ListPtrRecordsInvoker(request *model.ListPtrRecordsRequest) *ListPtrRecordsInvoker {
	requestDef := GenReqDefForListPtrRecords()
	return &ListPtrRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) RestorePtrRecord(request *model.RestorePtrRecordRequest) (*model.RestorePtrRecordResponse, error) {
	requestDef := GenReqDefForRestorePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestorePtrRecordResponse), nil
	}
}

func (c *DnsClient) RestorePtrRecordInvoker(request *model.RestorePtrRecordRequest) *RestorePtrRecordInvoker {
	requestDef := GenReqDefForRestorePtrRecord()
	return &RestorePtrRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowPtrRecordSet(request *model.ShowPtrRecordSetRequest) (*model.ShowPtrRecordSetResponse, error) {
	requestDef := GenReqDefForShowPtrRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPtrRecordSetResponse), nil
	}
}

func (c *DnsClient) ShowPtrRecordSetInvoker(request *model.ShowPtrRecordSetRequest) *ShowPtrRecordSetInvoker {
	requestDef := GenReqDefForShowPtrRecordSet()
	return &ShowPtrRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) UpdatePtrRecord(request *model.UpdatePtrRecordRequest) (*model.UpdatePtrRecordResponse, error) {
	requestDef := GenReqDefForUpdatePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePtrRecordResponse), nil
	}
}

func (c *DnsClient) UpdatePtrRecordInvoker(request *model.UpdatePtrRecordRequest) *UpdatePtrRecordInvoker {
	requestDef := GenReqDefForUpdatePtrRecord()
	return &UpdatePtrRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) CreateRecordSet(request *model.CreateRecordSetRequest) (*model.CreateRecordSetResponse, error) {
	requestDef := GenReqDefForCreateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetResponse), nil
	}
}

func (c *DnsClient) CreateRecordSetInvoker(request *model.CreateRecordSetRequest) *CreateRecordSetInvoker {
	requestDef := GenReqDefForCreateRecordSet()
	return &CreateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) CreateRecordSetWithLine(request *model.CreateRecordSetWithLineRequest) (*model.CreateRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForCreateRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetWithLineResponse), nil
	}
}

func (c *DnsClient) CreateRecordSetWithLineInvoker(request *model.CreateRecordSetWithLineRequest) *CreateRecordSetWithLineInvoker {
	requestDef := GenReqDefForCreateRecordSetWithLine()
	return &CreateRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) DeleteRecordSet(request *model.DeleteRecordSetRequest) (*model.DeleteRecordSetResponse, error) {
	requestDef := GenReqDefForDeleteRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetResponse), nil
	}
}

func (c *DnsClient) DeleteRecordSetInvoker(request *model.DeleteRecordSetRequest) *DeleteRecordSetInvoker {
	requestDef := GenReqDefForDeleteRecordSet()
	return &DeleteRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) DeleteRecordSets(request *model.DeleteRecordSetsRequest) (*model.DeleteRecordSetsResponse, error) {
	requestDef := GenReqDefForDeleteRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetsResponse), nil
	}
}

func (c *DnsClient) DeleteRecordSetsInvoker(request *model.DeleteRecordSetsRequest) *DeleteRecordSetsInvoker {
	requestDef := GenReqDefForDeleteRecordSets()
	return &DeleteRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListRecordSets(request *model.ListRecordSetsRequest) (*model.ListRecordSetsResponse, error) {
	requestDef := GenReqDefForListRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsResponse), nil
	}
}

func (c *DnsClient) ListRecordSetsInvoker(request *model.ListRecordSetsRequest) *ListRecordSetsInvoker {
	requestDef := GenReqDefForListRecordSets()
	return &ListRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListRecordSetsByZone(request *model.ListRecordSetsByZoneRequest) (*model.ListRecordSetsByZoneResponse, error) {
	requestDef := GenReqDefForListRecordSetsByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsByZoneResponse), nil
	}
}

func (c *DnsClient) ListRecordSetsByZoneInvoker(request *model.ListRecordSetsByZoneRequest) *ListRecordSetsByZoneInvoker {
	requestDef := GenReqDefForListRecordSetsByZone()
	return &ListRecordSetsByZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListRecordSetsWithLine(request *model.ListRecordSetsWithLineRequest) (*model.ListRecordSetsWithLineResponse, error) {
	requestDef := GenReqDefForListRecordSetsWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsWithLineResponse), nil
	}
}

func (c *DnsClient) ListRecordSetsWithLineInvoker(request *model.ListRecordSetsWithLineRequest) *ListRecordSetsWithLineInvoker {
	requestDef := GenReqDefForListRecordSetsWithLine()
	return &ListRecordSetsWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowRecordSet(request *model.ShowRecordSetRequest) (*model.ShowRecordSetResponse, error) {
	requestDef := GenReqDefForShowRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetResponse), nil
	}
}

func (c *DnsClient) ShowRecordSetInvoker(request *model.ShowRecordSetRequest) *ShowRecordSetInvoker {
	requestDef := GenReqDefForShowRecordSet()
	return &ShowRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowRecordSetByZone(request *model.ShowRecordSetByZoneRequest) (*model.ShowRecordSetByZoneResponse, error) {
	requestDef := GenReqDefForShowRecordSetByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetByZoneResponse), nil
	}
}

func (c *DnsClient) ShowRecordSetByZoneInvoker(request *model.ShowRecordSetByZoneRequest) *ShowRecordSetByZoneInvoker {
	requestDef := GenReqDefForShowRecordSetByZone()
	return &ShowRecordSetByZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowRecordSetWithLine(request *model.ShowRecordSetWithLineRequest) (*model.ShowRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForShowRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetWithLineResponse), nil
	}
}

func (c *DnsClient) ShowRecordSetWithLineInvoker(request *model.ShowRecordSetWithLineRequest) *ShowRecordSetWithLineInvoker {
	requestDef := GenReqDefForShowRecordSetWithLine()
	return &ShowRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) UpdateRecordSet(request *model.UpdateRecordSetRequest) (*model.UpdateRecordSetResponse, error) {
	requestDef := GenReqDefForUpdateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetResponse), nil
	}
}

func (c *DnsClient) UpdateRecordSetInvoker(request *model.UpdateRecordSetRequest) *UpdateRecordSetInvoker {
	requestDef := GenReqDefForUpdateRecordSet()
	return &UpdateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) UpdateRecordSets(request *model.UpdateRecordSetsRequest) (*model.UpdateRecordSetsResponse, error) {
	requestDef := GenReqDefForUpdateRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetsResponse), nil
	}
}

func (c *DnsClient) UpdateRecordSetsInvoker(request *model.UpdateRecordSetsRequest) *UpdateRecordSetsInvoker {
	requestDef := GenReqDefForUpdateRecordSets()
	return &UpdateRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) AssociateRouter(request *model.AssociateRouterRequest) (*model.AssociateRouterResponse, error) {
	requestDef := GenReqDefForAssociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouterResponse), nil
	}
}

func (c *DnsClient) AssociateRouterInvoker(request *model.AssociateRouterRequest) *AssociateRouterInvoker {
	requestDef := GenReqDefForAssociateRouter()
	return &AssociateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) CreatePrivateZone(request *model.CreatePrivateZoneRequest) (*model.CreatePrivateZoneResponse, error) {
	requestDef := GenReqDefForCreatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateZoneResponse), nil
	}
}

func (c *DnsClient) CreatePrivateZoneInvoker(request *model.CreatePrivateZoneRequest) *CreatePrivateZoneInvoker {
	requestDef := GenReqDefForCreatePrivateZone()
	return &CreatePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) CreatePublicZone(request *model.CreatePublicZoneRequest) (*model.CreatePublicZoneResponse, error) {
	requestDef := GenReqDefForCreatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicZoneResponse), nil
	}
}

func (c *DnsClient) CreatePublicZoneInvoker(request *model.CreatePublicZoneRequest) *CreatePublicZoneInvoker {
	requestDef := GenReqDefForCreatePublicZone()
	return &CreatePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) DeletePrivateZone(request *model.DeletePrivateZoneRequest) (*model.DeletePrivateZoneResponse, error) {
	requestDef := GenReqDefForDeletePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateZoneResponse), nil
	}
}

func (c *DnsClient) DeletePrivateZoneInvoker(request *model.DeletePrivateZoneRequest) *DeletePrivateZoneInvoker {
	requestDef := GenReqDefForDeletePrivateZone()
	return &DeletePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) DeletePublicZone(request *model.DeletePublicZoneRequest) (*model.DeletePublicZoneResponse, error) {
	requestDef := GenReqDefForDeletePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicZoneResponse), nil
	}
}

func (c *DnsClient) DeletePublicZoneInvoker(request *model.DeletePublicZoneRequest) *DeletePublicZoneInvoker {
	requestDef := GenReqDefForDeletePublicZone()
	return &DeletePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) DisassociateRouter(request *model.DisassociateRouterRequest) (*model.DisassociateRouterResponse, error) {
	requestDef := GenReqDefForDisassociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouterResponse), nil
	}
}

func (c *DnsClient) DisassociateRouterInvoker(request *model.DisassociateRouterRequest) *DisassociateRouterInvoker {
	requestDef := GenReqDefForDisassociateRouter()
	return &DisassociateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListPrivateZones(request *model.ListPrivateZonesRequest) (*model.ListPrivateZonesResponse, error) {
	requestDef := GenReqDefForListPrivateZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateZonesResponse), nil
	}
}

func (c *DnsClient) ListPrivateZonesInvoker(request *model.ListPrivateZonesRequest) *ListPrivateZonesInvoker {
	requestDef := GenReqDefForListPrivateZones()
	return &ListPrivateZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ListPublicZones(request *model.ListPublicZonesRequest) (*model.ListPublicZonesResponse, error) {
	requestDef := GenReqDefForListPublicZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicZonesResponse), nil
	}
}

func (c *DnsClient) ListPublicZonesInvoker(request *model.ListPublicZonesRequest) *ListPublicZonesInvoker {
	requestDef := GenReqDefForListPublicZones()
	return &ListPublicZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowPrivateZone(request *model.ShowPrivateZoneRequest) (*model.ShowPrivateZoneResponse, error) {
	requestDef := GenReqDefForShowPrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneResponse), nil
	}
}

func (c *DnsClient) ShowPrivateZoneInvoker(request *model.ShowPrivateZoneRequest) *ShowPrivateZoneInvoker {
	requestDef := GenReqDefForShowPrivateZone()
	return &ShowPrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowPrivateZoneNameServer(request *model.ShowPrivateZoneNameServerRequest) (*model.ShowPrivateZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPrivateZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

func (c *DnsClient) ShowPrivateZoneNameServerInvoker(request *model.ShowPrivateZoneNameServerRequest) *ShowPrivateZoneNameServerInvoker {
	requestDef := GenReqDefForShowPrivateZoneNameServer()
	return &ShowPrivateZoneNameServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowPublicZone(request *model.ShowPublicZoneRequest) (*model.ShowPublicZoneResponse, error) {
	requestDef := GenReqDefForShowPublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneResponse), nil
	}
}

func (c *DnsClient) ShowPublicZoneInvoker(request *model.ShowPublicZoneRequest) *ShowPublicZoneInvoker {
	requestDef := GenReqDefForShowPublicZone()
	return &ShowPublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) ShowPublicZoneNameServer(request *model.ShowPublicZoneNameServerRequest) (*model.ShowPublicZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPublicZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

func (c *DnsClient) ShowPublicZoneNameServerInvoker(request *model.ShowPublicZoneNameServerRequest) *ShowPublicZoneNameServerInvoker {
	requestDef := GenReqDefForShowPublicZoneNameServer()
	return &ShowPublicZoneNameServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) UpdatePrivateZone(request *model.UpdatePrivateZoneRequest) (*model.UpdatePrivateZoneResponse, error) {
	requestDef := GenReqDefForUpdatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateZoneResponse), nil
	}
}

func (c *DnsClient) UpdatePrivateZoneInvoker(request *model.UpdatePrivateZoneRequest) *UpdatePrivateZoneInvoker {
	requestDef := GenReqDefForUpdatePrivateZone()
	return &UpdatePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *DnsClient) UpdatePublicZone(request *model.UpdatePublicZoneRequest) (*model.UpdatePublicZoneResponse, error) {
	requestDef := GenReqDefForUpdatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneResponse), nil
	}
}

func (c *DnsClient) UpdatePublicZoneInvoker(request *model.UpdatePublicZoneRequest) *UpdatePublicZoneInvoker {
	requestDef := GenReqDefForUpdatePublicZone()
	return &UpdatePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
