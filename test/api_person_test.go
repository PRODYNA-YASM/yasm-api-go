/*
YASM (Yet Another Skill Management) API

Testing PersonAPIService

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

func Test_client_PersonAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PersonAPIService AddPersonCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var certificationId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonCertification(context.Background(), personId, certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddPersonInterest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonInterest(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddPersonLanguage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var languageId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonLanguage(context.Background(), personId, languageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddPersonOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var officeId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonOffice(context.Background(), personId, officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddPersonProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var profileId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonProfile(context.Background(), personId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddPersonSkillExperience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonSkillExperience(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddPersonSkillExperiences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.AddPersonSkillExperiences(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PersonAPI.AddProjectParticipation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService AddSkillConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.PersonAPI.AddSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService CreateAvailability", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.CreateAvailability(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService CreatePerson", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PersonAPI.CreatePerson(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeleteAvailability", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var availabilityId string

		resp, httpRes, err := apiClient.PersonAPI.DeleteAvailability(context.Background(), personId, availabilityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePerson", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePerson(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePersonCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var certificationId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePersonCertification(context.Background(), personId, certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePersonInterest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePersonInterest(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePersonOffice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var officeId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePersonOffice(context.Background(), personId, officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePersonProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var profileId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePersonProfile(context.Background(), personId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePersonSkillExperience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePersonSkillExperience(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeletePersonSkillExperiences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.DeletePersonSkillExperiences(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService DeleteProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string

		resp, httpRes, err := apiClient.PersonAPI.DeleteProjectParticipation(context.Background(), projectParticipationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService GeneratePersonProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.GeneratePersonProfile(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService GetAllBusinessDepartments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PersonAPI.GetAllBusinessDepartments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService GetAvailabilities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.GetAvailabilities(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService GetPerson", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.GetPerson(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService ReadPersonPicture", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.ReadPersonPicture(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService ReadPersonProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.ReadPersonProjectParticipation(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService ReadProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string

		resp, httpRes, err := apiClient.PersonAPI.ReadProjectParticipation(context.Background(), projectParticipationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService RemovePersonLanguage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var languageId string

		resp, httpRes, err := apiClient.PersonAPI.RemovePersonLanguage(context.Background(), personId, languageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService RemoveSkillConfirmation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string
		var skillId string
		var confirmingPersonId string

		resp, httpRes, err := apiClient.PersonAPI.RemoveSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService SearchPersons", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PersonAPI.SearchPersons(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdateAvailability", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var availabilityId string

		resp, httpRes, err := apiClient.PersonAPI.UpdateAvailability(context.Background(), personId, availabilityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdatePerson", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.UpdatePerson(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdatePersonCertification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var certificationId string

		resp, httpRes, err := apiClient.PersonAPI.UpdatePersonCertification(context.Background(), personId, certificationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdatePersonPicture", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.UpdatePersonPicture(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdatePersonSkillExperience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var skillId string

		resp, httpRes, err := apiClient.PersonAPI.UpdatePersonSkillExperience(context.Background(), personId, skillId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdatePersonSkillExperiences", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string

		resp, httpRes, err := apiClient.PersonAPI.UpdatePersonSkillExperiences(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UpdateProjectParticipation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectParticipationId string

		resp, httpRes, err := apiClient.PersonAPI.UpdateProjectParticipation(context.Background(), projectParticipationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService UupdatePersonLanguage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId string
		var languageId string

		resp, httpRes, err := apiClient.PersonAPI.UupdatePersonLanguage(context.Background(), personId, languageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
