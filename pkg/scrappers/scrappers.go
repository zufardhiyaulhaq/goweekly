package scrappers

import (
	communityv1alpha1 "github.com/zufardhiyaulhaq/community-operator-v2/api/v1alpha1"
)

type Scrapper interface {
	GetName() (string, error)
	GetArticles() ([]communityv1alpha1.WeeklySpec_Article, error)
}
