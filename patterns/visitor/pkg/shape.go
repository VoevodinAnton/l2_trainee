package pkg

type shape interface {
	GetType() string
	Accept(visitor)
}
