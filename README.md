# goweekly

Get data from golang news and create Weekly CRDs based on community-operator and push to git datastore

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github master branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/goweekly/Master)](https://github.com/zufardhiyaulhaq/goweekly/actions/workflows/master.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/goweekly)](https://github.com/zufardhiyaulhaq/goweekly/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/goweekly)](https://github.com/zufardhiyaulhaq/goweekly/pulls)

[![Bugs](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=bugs)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)
[![Code Smells](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=code_smells)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)
[![Coverage](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=coverage)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)
[![Maintainability Rating](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=sqale_rating)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)
[![Quality Gate Status](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=alert_status)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)
[![Reliability Rating](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=reliability_rating)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)
[![Security Rating](https://sonarqube.zufardhiyaulhaq.com/api/project_badges/measure?project=goweekly&metric=security_rating)](https://sonarqube.zufardhiyaulhaq.com/dashboard?id=goweekly)

## Installing the Chart

To install the chart with the release name `my-release` and secret `foo`:

```console
kubectl apply -f secret.yaml

helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install goweekly zufardhiyaulhaq/goweekly --values values.yaml --set secret=foo
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cronSchedule | string | `"0 8 * * 0"` |  |
| image.repository | string | `"zufardhiyaulhaq/goweekly"` |  |
| image.tag | string | `"v1.0.0"` |  |
| secret | string | `""` |  |
| weekly.community | string | `"Golang Indonesia Community"` |  |
| weekly.image_url | string | `"https://trungtq.com/wp-content/uploads/2018/12/GO-3.png"` |  |
| weekly.namespace | string | `"golang-community"` |  |
| weekly.tags | string | `"weekly,golang"` |  |

