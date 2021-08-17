package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zufardhiyaulhaq/goweekly/pkg/models"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

func TestWeeklyToYaml(t *testing.T) {
	weekly := models.Weekly{
		APIVersion: "community.io/v1alpha1",
		Kind:       "Weekly",
		Metadata: models.WeeklyMetadata{
			Name:      "foo",
			Namespace: "bar",
		},
		Spec: communityv1alpha1.WeeklySpec{
			Name:      "foo",
			Date:      "bar",
			Image:     "foo",
			Community: "bar",
			Tags: []string{
				"foo",
				"bar",
			},
			Articles: []communityv1alpha1.ArticleSpec{
				{
					Title: "foo",
					Url:   "bar",
					Type:  "foo",
				},
				{
					Title: "foo",
					Url:   "bar",
					Type:  "foo",
				},
			},
		},
	}

	weeklyBytes, err := weekly.ToYaml()
	assert.NoError(t, err)

	expectedOutput := `apiVersion: community.io/v1alpha1
kind: Weekly
metadata:
  name: foo
  namespace: bar
spec:
  name: foo
  date: bar
  image: foo
  community: bar
  tags:
  - foo
  - bar
  articles:
  - title: foo
    url: bar
    type: foo
  - title: foo
    url: bar
    type: foo
`
	assert.Equal(t, expectedOutput, string(weeklyBytes))
}
