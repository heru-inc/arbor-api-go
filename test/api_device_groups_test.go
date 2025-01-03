/*
arborxr-api-v2

Testing DeviceGroupsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package arborapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/heru-inc/arbor-api-go"
)

func Test_arborapi_DeviceGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceGroupsAPIService DeviceGroupsAddReleaseChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceGroupId string

		httpRes, err := apiClient.DeviceGroupsAPI.DeviceGroupsAddReleaseChannel(context.Background(), deviceGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceGroupsAPIService DeviceGroupsGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeviceGroupsAPI.DeviceGroupsGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceGroupsAPIService DeviceGroupsRemoveReleaseChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceGroupId string

		httpRes, err := apiClient.DeviceGroupsAPI.DeviceGroupsRemoveReleaseChannel(context.Background(), deviceGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
