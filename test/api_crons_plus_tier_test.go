/*
LangGraph Platform

Testing CronsPlusTierAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package langgraphgo

import (
	"context"
	"testing"

	openapiclient "github.com/chelala/langgraphgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_langgraphgo_CronsPlusTierAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CronsPlusTierAPIService CreateCronRunsCronsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CronsPlusTierAPI.CreateCronRunsCronsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronsPlusTierAPIService CreateThreadCronThreadsThreadIdRunsCronsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var threadId string

		resp, httpRes, err := apiClient.CronsPlusTierAPI.CreateThreadCronThreadsThreadIdRunsCronsPost(context.Background(), threadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronsPlusTierAPIService DeleteCronRunsCronsCronIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cronId string

		resp, httpRes, err := apiClient.CronsPlusTierAPI.DeleteCronRunsCronsCronIdDelete(context.Background(), cronId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronsPlusTierAPIService SearchCronsRunsCronsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CronsPlusTierAPI.SearchCronsRunsCronsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
