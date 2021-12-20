package fetcher

import "fmt"

type CardError struct {
	number int
	set    string
	err    error
}

func NewCardError(number int, set string, err error) *CardError {
	return &CardError{
		number: number,
		set:    set,
		err:    err,
	}
}

func (ce *CardError) Error() string {
	return fmt.Errorf("card #%d from set %q failed with %w", ce.number, ce.set, ce.err).Error()
}
