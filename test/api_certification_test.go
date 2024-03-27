/*
YASM (Yet Another Skill Management) API

Testing CertificationAPIService

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

func Test_client_CertificationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificationAPIService AddPersonCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var certificationId string

		resp, httpRes, err := apiClient.CertificationAPI.AddPersonCertification(context.Background(), personId, certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService AddSkillToCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var skillId string

		resp, httpRes, err := apiClient.CertificationAPI.AddSkillToCertification(context.Background(), certificationId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService CreateCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificationAPI.CreateCertification(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService DeleteCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string

		resp, httpRes, err := apiClient.CertificationAPI.DeleteCertification(context.Background(), certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService DeletePersonCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var certificationId string

		resp, httpRes, err := apiClient.CertificationAPI.DeletePersonCertification(context.Background(), personId, certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService DeleteSkillFromCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var skillId string

		resp, httpRes, err := apiClient.CertificationAPI.DeleteSkillFromCertification(context.Background(), certificationId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService GetCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string

		resp, httpRes, err := apiClient.CertificationAPI.GetCertification(context.Background(), certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService MoveCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var organizationId string

		resp, httpRes, err := apiClient.CertificationAPI.MoveCertification(context.Background(), certificationId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService SearchCertifications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificationAPI.SearchCertifications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService UpdateCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string

		resp, httpRes, err := apiClient.CertificationAPI.UpdateCertification(context.Background(), certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService UpdatePersonCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var certificationId string

		resp, httpRes, err := apiClient.CertificationAPI.UpdatePersonCertification(context.Background(), personId, certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationAPIService UpdateSkillInCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var skillId string

		resp, httpRes, err := apiClient.CertificationAPI.UpdateSkillInCertification(context.Background(), certificationId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
