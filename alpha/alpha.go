package alpha

type Int *int64
type String *string

type error *s

/*
func (e *Error) Error() string {
	if e != nil {
		return string(**e)
	}
	return ""
}

var _ error = Error(nil)
*/

type s struct {
	s string
}
type i struct {
	i int64
}
