package partial

import "encoding/json"

type PartialJSON[E any] struct {
	json.RawMessage
}

func(it PartialJSON[E]) Get() (*E, error) {
	var e E

	if err := it.Apply(&e); err != nil {
		return nil, err
	}

	return &e, nil
}

func(it PartialJSON[E]) Apply(e *E) error {
	return json.Unmarshal(it.RawMessage, e)
}
