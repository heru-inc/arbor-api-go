/*
arborxr-api-v2

Testing DevicesAPIService

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

func Test_arborapi_DevicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevicesAPIService DevicesAddReleaseChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DevicesAPI.DevicesAddReleaseChannel(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesCheckFingerprint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DevicesAPI.DevicesCheckFingerprint(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DevicesAPI.DevicesDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DevicesAPI.DevicesDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesLaunchApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string
		var appId string

		httpRes, err := apiClient.DevicesAPI.DevicesLaunchApp(context.Background(), deviceId, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesReboot", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DevicesAPI.DevicesReboot(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesRemoveReleaseChannel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DevicesAPI.DevicesRemoveReleaseChannel(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DevicesAPIService DevicesUpdateDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DevicesAPI.DevicesUpdateDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
