/*
YASM (Yet Another Skill Management) API

Testing OfficeAPIService

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

func Test_client_OfficeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OfficeAPIService AddOrganizationOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.AddOrganizationOffice(context.Background(), organizationId, officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService CreateOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OfficeAPI.CreateOffice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService DeleteOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.DeleteOffice(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService DeleteOfficeCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string
		var countryId string

		resp, httpRes, err := apiClient.OfficeAPI.DeleteOfficeCountry(context.Background(), officeId, countryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService DeleteOrganizationOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.DeleteOrganizationOffice(context.Background(), organizationId, officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService DeletePersonOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.DeletePersonOffice(context.Background(), personId, officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService GetOfficeDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.GetOfficeDetails(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService SearchOffices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OfficeAPI.SearchOffices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService UpdateOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.UpdateOffice(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService UpdateOfficeCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string
		var countryId string

		resp, httpRes, err := apiClient.OfficeAPI.UpdateOfficeCountry(context.Background(), officeId, countryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OfficeAPIService UpdatePersonOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var officeId string

		resp, httpRes, err := apiClient.OfficeAPI.UpdatePersonOffice(context.Background(), personId, officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
