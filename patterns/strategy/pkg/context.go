package pkg

type privateValues struct {
	actionStrategy Strategy
	UserFilters    map[string]int
}

func InitStrategy(s Strategy) *privateValues {
	f := make(map[string]int)

	return &privateValues{
		actionStrategy: s,
		UserFilters:    f,
	}
}

func (v *privateValues) SetStrategy(s Strategy) {
	v.actionStrategy = s
}

func GetData(v *privateValues) {
	v.actionStrategy.doSearch(v.UserFilters)
}
