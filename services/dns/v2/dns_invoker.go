package v2

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/invoker"
	"github.com/fecloud-sdk/fecloud-sdk-go/services/dns/v2/model"
)

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ListNameServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNameServersInvoker) Invoke() (*model.ListNameServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNameServersResponse), nil
	}
}

type ShowApiInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiInfoInvoker) Invoke() (*model.ShowApiInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiInfoResponse), nil
	}
}

type ShowDomainQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainQuotaInvoker) Invoke() (*model.ShowDomainQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainQuotaResponse), nil
	}
}

type CreateEipRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEipRecordSetInvoker) Invoke() (*model.CreateEipRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEipRecordSetResponse), nil
	}
}

type ListPtrRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPtrRecordsInvoker) Invoke() (*model.ListPtrRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPtrRecordsResponse), nil
	}
}

type RestorePtrRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestorePtrRecordInvoker) Invoke() (*model.RestorePtrRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestorePtrRecordResponse), nil
	}
}

type ShowPtrRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPtrRecordSetInvoker) Invoke() (*model.ShowPtrRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPtrRecordSetResponse), nil
	}
}

type UpdatePtrRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePtrRecordInvoker) Invoke() (*model.UpdatePtrRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePtrRecordResponse), nil
	}
}

type CreateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetInvoker) Invoke() (*model.CreateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetResponse), nil
	}
}

type CreateRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetWithLineInvoker) Invoke() (*model.CreateRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetWithLineResponse), nil
	}
}

type DeleteRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordSetInvoker) Invoke() (*model.DeleteRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordSetResponse), nil
	}
}

type DeleteRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordSetsInvoker) Invoke() (*model.DeleteRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordSetsResponse), nil
	}
}

type ListRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsInvoker) Invoke() (*model.ListRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsResponse), nil
	}
}

type ListRecordSetsByZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsByZoneInvoker) Invoke() (*model.ListRecordSetsByZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsByZoneResponse), nil
	}
}

type ListRecordSetsWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsWithLineInvoker) Invoke() (*model.ListRecordSetsWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsWithLineResponse), nil
	}
}

type ShowRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetInvoker) Invoke() (*model.ShowRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetResponse), nil
	}
}

type ShowRecordSetByZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetByZoneInvoker) Invoke() (*model.ShowRecordSetByZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetByZoneResponse), nil
	}
}

type ShowRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetWithLineInvoker) Invoke() (*model.ShowRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetWithLineResponse), nil
	}
}

type UpdateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordSetInvoker) Invoke() (*model.UpdateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordSetResponse), nil
	}
}

type UpdateRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordSetsInvoker) Invoke() (*model.UpdateRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordSetsResponse), nil
	}
}

type AssociateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRouterInvoker) Invoke() (*model.AssociateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRouterResponse), nil
	}
}

type CreatePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateZoneInvoker) Invoke() (*model.CreatePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateZoneResponse), nil
	}
}

type CreatePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublicZoneInvoker) Invoke() (*model.CreatePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublicZoneResponse), nil
	}
}

type DeletePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateZoneInvoker) Invoke() (*model.DeletePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateZoneResponse), nil
	}
}

type DeletePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublicZoneInvoker) Invoke() (*model.DeletePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublicZoneResponse), nil
	}
}

type DisassociateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateRouterInvoker) Invoke() (*model.DisassociateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateRouterResponse), nil
	}
}

type ListPrivateZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateZonesInvoker) Invoke() (*model.ListPrivateZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateZonesResponse), nil
	}
}

type ListPublicZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicZonesInvoker) Invoke() (*model.ListPublicZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicZonesResponse), nil
	}
}

type ShowPrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateZoneInvoker) Invoke() (*model.ShowPrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateZoneResponse), nil
	}
}

type ShowPrivateZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateZoneNameServerInvoker) Invoke() (*model.ShowPrivateZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

type ShowPublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicZoneInvoker) Invoke() (*model.ShowPublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicZoneResponse), nil
	}
}

type ShowPublicZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicZoneNameServerInvoker) Invoke() (*model.ShowPublicZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

type UpdatePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateZoneInvoker) Invoke() (*model.UpdatePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateZoneResponse), nil
	}
}

type UpdatePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicZoneInvoker) Invoke() (*model.UpdatePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicZoneResponse), nil
	}
}
