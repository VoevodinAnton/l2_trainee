package pkg

type Strategy interface {
	doSearch(filters map[string]int)
}
