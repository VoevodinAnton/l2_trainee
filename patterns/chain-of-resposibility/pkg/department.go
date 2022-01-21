package pkg

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
