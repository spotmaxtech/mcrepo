package pkg

var CurrentRegistry Registry

type Registry interface {
	ListRepo()
	ListRepoTag(repoName string)
}
