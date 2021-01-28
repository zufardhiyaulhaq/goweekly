## Development
For small things like fixing typos in documentation, you can [make edits through GitHub](https://help.github.com/articles/editing-files-in-another-user-s-repository/), which will handle forking and making a pull request (PR) for you. For anything bigger or more complex, you'll probably want to set up a development environment on your machine, a quick procedure for which is as folows:

### Setup your machine
Prerequisites:
- make
- [Go 1.13+](https://golang.org/doc/install)

- export variable for developing
```
### Github datastore information
### Will store Weekly object CRD as a yaml file

# Github token
export GITHUB_TOKEN="token"

# Github username or organization
export GITHUB_ORGANIZATION="zufardhiyaulhaq"

# Github Repository to store the weekly content
export GITHUB_REPOSITORY="community-ops"

# Github path in the repository 
export GITHUB_REPOSITORY_PATH="./manifest/golang-community/"

# Github Branch
export GITHUB_BRANCH="master"

### Weekly CRD specification

# Community name for the weekly
export COMMUNITY="Golang Indonesia Community"

# Tags for the weekly
export TAGS="weekly,golang"

# Namespace for weekly
export NAMESPACE="golang-community"

# Image URL for weekly
export IMAGE="https://trungtq.com/wp-content/uploads/2018/12/GO-3.png"
```
- Build & Run
```
go build -o goweekly cmd/goweekly/*.go
./goweekly
```

### Build Docker Image
- Build image
```
make build REPOSITORY=username/repository TAG=tag
```
