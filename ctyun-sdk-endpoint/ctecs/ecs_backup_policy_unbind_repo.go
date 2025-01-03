package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyUnbindRepoApi
type EcsBackupPolicyUnbindRepoApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupPolicyUnbindRepoApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyUnbindRepoApi {
	return &EcsBackupPolicyUnbindRepoApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/unbind-repo",
		},
	}
}

func (this *EcsBackupPolicyUnbindRepoApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyUnbindRepoRequest) (*EcsBackupPolicyUnbindRepoResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupPolicyUnbindRepoRealRequest{
		RegionID: req.RegionID,
		PolicyID: req.PolicyID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupPolicyUnbindRepoRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupPolicyUnbindRepoResponse{
		PolicyID: realResponse.PolicyID,
	}, nil
}

type EcsBackupPolicyUnbindRepoRealRequest struct {
	RegionID *string `json:"regionID,omitempty"`
	PolicyID *string `json:"policyID,omitempty"`
}

type EcsBackupPolicyUnbindRepoRequest struct {
	RegionID *string
	PolicyID *string
}

type EcsBackupPolicyUnbindRepoRealResponse struct {
	PolicyID string `json:"policyID,omitempty"`
}

type EcsBackupPolicyUnbindRepoResponse struct {
	PolicyID string
}
