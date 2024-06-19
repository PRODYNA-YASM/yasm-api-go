/*
YASM (Yet Another Skill Management) API

Testing SkillAPIService

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

func Test_client_SkillAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SkillAPIService AddPersonInterest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.AddPersonInterest(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService AddPersonSkillExperience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.AddPersonSkillExperience(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService AddPersonSkillExperiences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.SkillAPI.AddPersonSkillExperiences(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService AddSkillConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.SkillAPI.AddSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService AddSkillToCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.AddSkillToCertification(context.Background(), certificationId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService AddSkillToParentSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var skillId string
		var parentSkillId string

		resp, httpRes, err := apiClient.SkillAPI.AddSkillToParentSkill(context.Background(), skillId, parentSkillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService CreateSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SkillAPI.CreateSkill(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService DeletePersonInterest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.DeletePersonInterest(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService DeletePersonSkillExperience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.DeletePersonSkillExperience(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService DeletePersonSkillExperiences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.SkillAPI.DeletePersonSkillExperiences(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService DeleteSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.DeleteSkill(context.Background(), skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService DeleteSkillFromCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.DeleteSkillFromCertification(context.Background(), certificationId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService GetSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.GetSkill(context.Background(), skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService ReadPersonSkillStatistics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.ReadPersonSkillStatistics(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService RemoveSkillConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.SkillAPI.RemoveSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService RemoveSkillFromParentSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var skillId string
		var parentSkillId string

		resp, httpRes, err := apiClient.SkillAPI.RemoveSkillFromParentSkill(context.Background(), skillId, parentSkillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService SearchSkills", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SkillAPI.SearchSkills(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService UpdatePersonSkillExperience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.UpdatePersonSkillExperience(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService UpdatePersonSkillExperiences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.SkillAPI.UpdatePersonSkillExperiences(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService UpdateSkill", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.UpdateSkill(context.Background(), skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SkillAPIService UpdateSkillInCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certificationId string
		var skillId string

		resp, httpRes, err := apiClient.SkillAPI.UpdateSkillInCertification(context.Background(), certificationId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
