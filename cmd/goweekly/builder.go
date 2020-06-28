package main

import (
	"os"
	"strings"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
)

type Builder struct {
	APIVersion string                       `json:"apiVersion"`
	Kind       string                       `json:"kind"`
	Metadata   BuilderMetadata              `json:"metadata"`
	Spec       communityv1alpha1.WeeklySpec `json:"spec"`
}

type BuilderMetadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (b *Builder) build(name string, articles []communityv1alpha1.ArticleSpec) {
	b.APIVersion = "community.io/v1alpha1"
	b.Kind = "Weekly"
	b.Metadata.Name = name
	b.Metadata.Namespace = os.Getenv("NAMESPACE")

	b.Spec.Name = name
	b.Spec.Community = os.Getenv("COMMUNITY")
	b.Spec.Tags = strings.Split(os.Getenv("TAGS"), ",")
	b.Spec.Date = GetDate()
	b.Spec.Image = os.Getenv("IMAGE")
	b.Spec.Articles = articles
}
