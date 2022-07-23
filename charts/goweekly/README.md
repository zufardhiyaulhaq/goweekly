# goweekly

Get data from golang news and create Weekly CRDs based on community-operator and push to git datastore

![Version: 2.0.0](https://img.shields.io/badge/Version-2.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.0.0](https://img.shields.io/badge/AppVersion-2.0.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github master branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/goweekly/Master)](https://github.com/zufardhiyaulhaq/goweekly/actions/workflows/master.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/goweekly)](https://github.com/zufardhiyaulhaq/goweekly/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/goweekly)](https://github.com/zufardhiyaulhaq/goweekly/pulls)

## Installing the Chart

To install the chart with the release name `my-goweekly` and secret `foo`:

```console
kubectl apply -f secret.yaml

helm repo add goweekly https://zufardhiyaulhaq.com/goweekly/charts/releases/
helm install my-goweekly goweekly/goweekly --values values.yaml
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cronSchedule | string | `"0 8 * * 0"` |  |
| image.repository | string | `"zufardhiyaulhaq/goweekly"` |  |
| image.tag | string | `"v2.0.0"` |  |
| secret | string | `""` |  |
| weekly.community | string | `"golang"` |  |
| weekly.image_url | string | `"https://trungtq.com/wp-content/uploads/2018/12/GO-3.png"` |  |
| weekly.namespace | string | `"community"` |  |
| weekly.tags | string | `"weekly,golang"` |  |

