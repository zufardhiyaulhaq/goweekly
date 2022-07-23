package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zufardhiyaulhaq/goweekly/pkg/models"

	communityv1alpha1 "github.com/zufardhiyaulhaq/community-operator-v2/api/v1alpha1"
)

func TestWeeklyToYaml(t *testing.T) {
	weekly := models.Weekly{
		APIVersion: "community.zufardhiyaulhaq.com/v1alpha1",
		Kind:       "Weekly",
		Metadata: models.WeeklyMetadata{
			Name:      "foo",
			Namespace: "bar",
		},
		Spec: communityv1alpha1.WeeklySpec{
			Community: []string{
				"foo",
			},
			Spec: communityv1alpha1.WeeklySpec_Spec{
				Name:     "foo",
				Date:     "bar",
				ImageUrl: "foo",
				Tags: []string{
					"foo",
					"bar",
				},
				Articles: []communityv1alpha1.WeeklySpec_Article{
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
		},
	}

	weeklyBytes, err := weekly.ToYaml()
	assert.NoError(t, err)

	expectedOutput := `apiVersion: community.zufardhiyaulhaq.com/v1alpha1
kind: Weekly
metadata:
  name: foo
  namespace: bar
spec:
  community:
  - foo
  spec:
    name: foo
    date: bar
    imageUrl: foo
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
