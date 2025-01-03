package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupRepoRenewApi 续订云主机备份存储库
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=6908&data=87&isNormal=1
type EcsBackupRepoRenewApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupRepoRenewApi(client *ctyunsdk.CtyunClient) *EcsBackupRepoRenewApi {
	return &EcsBackupRepoRenewApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-repo/renew",
		},
	}
}

func (this *EcsBackupRepoRenewApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupRepoRenewRequest) (*EcsBackupRepoRenewResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupRepoRenewRealRequest{
		RegionID:        req.RegionID,
		RepositoryID:    req.RepositoryID,
		CycleCount:      req.CycleCount,
		CycleType:       req.CycleType,
		ClientToken:     req.ClientToken,
		PayVoucherPrice: req.PayVoucherPrice,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupRepoRenewRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupRepoRenewResponse{
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
		RegionID:      realResponse.RegionID,
	}, nil
}

type EcsBackupRepoRenewRealRequest struct {
	RegionID        *string  `json:"regionID,omitempty"`
	RepositoryID    *string  `json:"repositoryID,omitempty"`
	CycleCount      *int     `json:"cycleCount,omitempty"`
	CycleType       *string  `json:"cycleType,omitempty"`
	ClientToken     *string  `json:"clientToken,omitempty"`
	PayVoucherPrice *float64 `json:"payVoucherPrice,omitempty"`
}

type EcsBackupRepoRenewRequest struct {
	RegionID        *string
	RepositoryID    *string
	CycleCount      *int
	CycleType       *string
	ClientToken     *string
	PayVoucherPrice *float64
}

type EcsBackupRepoRenewRealResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
	RegionID      string `json:"regionID,omitempty"`
}

type EcsBackupRepoRenewResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
