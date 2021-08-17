package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zufardhiyaulhaq/goweekly/pkg/models"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

func TestWeeklyFilenameBuilder(t *testing.T) {
	filename, err := models.NewWeeklyFilenameBuilder().
		SetName("foo").
		Build()

	assert.NoError(t, err)
	assert.Equal(t, "foo.yaml", filename)
}

func TestWeeklySpecBuilder(t *testing.T) {
	spec, err := models.NewWeeklySpecBuilder().
		SetName("foo").
		SetDate("bar").
		SetImage("foo").
		SetCommunity("bar").
		SetTags([]string{
			"foo",
			"bar",
		}).
		SetArticles([]communityv1alpha1.ArticleSpec{
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
		).Build()
	assert.NoError(t, err)
	assert.Equal(t, communityv1alpha1.WeeklySpec{
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
	}, spec)
}

func TestWeeklyBuilder(t *testing.T) {
	weekly, err := models.NewWeeklyBuilder().
		SetName("foo").
		SetNamespace("bar").
		SetSpec(communityv1alpha1.WeeklySpec{
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
		}).
		Build()

	assert.NoError(t, err)
	assert.Equal(t, "foo", weekly.Metadata.Name)
	assert.Equal(t, "bar", weekly.Metadata.Namespace)
}
