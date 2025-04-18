/*
YASM (Yet Another Skill Management) API

Testing ProjectImageAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func Test_client_ProjectImageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectImageAPIService CreateProjectImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectImageAPI.CreateProjectImage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectImageAPIService DeleteProjectImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectImageId string

		resp, httpRes, err := apiClient.ProjectImageAPI.DeleteProjectImage(context.Background(), projectImageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectImageAPIService ReadProjectImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectImageId string

		resp, httpRes, err := apiClient.ProjectImageAPI.ReadProjectImage(context.Background(), projectImageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectImageAPIService SearchProjectImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectImageAPI.SearchProjectImages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
