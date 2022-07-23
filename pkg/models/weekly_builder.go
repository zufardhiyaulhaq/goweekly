package models

import (
	"strings"

	communityv1alpha1 "github.com/zufardhiyaulhaq/community-operator-v2/api/v1alpha1"
)

type WeeklySpecBuilder struct {
	Spec communityv1alpha1.WeeklySpec
}

func (ws WeeklySpecBuilder) SetName(name string) WeeklySpecBuilder {
	ws.Spec.Spec.Name = strings.ReplaceAll(name, "#", "")
	return ws
}

func (ws WeeklySpecBuilder) SetDate(date string) WeeklySpecBuilder {
	ws.Spec.Spec.Date = date
	return ws
}

func (ws WeeklySpecBuilder) SetImageUrl(imageUrl string) WeeklySpecBuilder {
	ws.Spec.Spec.ImageUrl = imageUrl
	return ws
}

func (ws WeeklySpecBuilder) SetCommunity(community []string) WeeklySpecBuilder {
	ws.Spec.Community = community
	return ws
}

func (ws WeeklySpecBuilder) SetTags(tags []string) WeeklySpecBuilder {
	ws.Spec.Spec.Tags = tags
	return ws
}

func (ws WeeklySpecBuilder) SetArticles(articles []communityv1alpha1.WeeklySpec_Article) WeeklySpecBuilder {
	ws.Spec.Spec.Articles = articles
	return ws
}

func (ws WeeklySpecBuilder) Build() (communityv1alpha1.WeeklySpec, error) {
	return ws.Spec, nil
}

func NewWeeklySpecBuilder() WeeklySpecBuilder {
	return WeeklySpecBuilder{}
}

type WeeklyBuilder struct {
	Weekly Weekly
}

func (w WeeklyBuilder) SetName(name string) WeeklyBuilder {
	w.Weekly.Metadata.Name = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name, "#", ""), " ", "-"))
	return w
}

func (w WeeklyBuilder) SetNamespace(namespace string) WeeklyBuilder {
	w.Weekly.Metadata.Namespace = namespace
	return w
}

func (w WeeklyBuilder) SetSpec(spec communityv1alpha1.WeeklySpec) WeeklyBuilder {
	w.Weekly.Spec = spec
	return w
}

func (w WeeklyBuilder) Build() (Weekly, error) {
	return w.Weekly, nil
}

func NewWeeklyBuilder() WeeklyBuilder {
	return WeeklyBuilder{
		Weekly: Weekly{
			APIVersion: "community.zufardhiyaulhaq.com/v1alpha1",
			Kind:       "Weekly",
		},
	}
}

type WeeklyFilenameBuilder struct {
	Name string
}

func (wf WeeklyFilenameBuilder) SetName(name string) WeeklyFilenameBuilder {
	wf.Name = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name, "#", ""), " ", "-"))
	return wf
}

func (w WeeklyFilenameBuilder) Build() (string, error) {
	return w.Name + ".yaml", nil
}

func NewWeeklyFilenameBuilder() WeeklyFilenameBuilder {
	return WeeklyFilenameBuilder{}
}
