package partial

type Partial[E any] interface {
	Apply(e *E) error
}
