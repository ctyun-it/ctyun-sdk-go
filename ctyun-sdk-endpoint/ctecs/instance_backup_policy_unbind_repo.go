package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// InstanceBackupPolicyUnbindRepoApi
type InstanceBackupPolicyUnbindRepoApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewInstanceBackupPolicyUnbindRepoApi(client *ctyunsdk.CtyunClient) *InstanceBackupPolicyUnbindRepoApi {
	return &InstanceBackupPolicyUnbindRepoApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/unbind-repo",
		},
	}
}

func (this *InstanceBackupPolicyUnbindRepoApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *InstanceBackupPolicyUnbindRepoRequest) (*InstanceBackupPolicyUnbindRepoResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&InstanceBackupPolicyUnbindRepoRealRequest{
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

	var realResponse InstanceBackupPolicyUnbindRepoRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &InstanceBackupPolicyUnbindRepoResponse{
		PolicyID: realResponse.PolicyID,
	}, nil
}

type InstanceBackupPolicyUnbindRepoRealRequest struct {
	RegionID string `json:"regionID,omitempty"`
	PolicyID string `json:"policyID,omitempty"`
}

type InstanceBackupPolicyUnbindRepoRequest struct {
	RegionID string
	PolicyID string
}

type InstanceBackupPolicyUnbindRepoRealResponse struct {
	PolicyID string `json:"policyID,omitempty"`
}

type InstanceBackupPolicyUnbindRepoResponse struct {
	PolicyID string
}
