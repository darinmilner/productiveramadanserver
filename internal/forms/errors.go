package forms

//errors type map of slice of strings
type errors map[string][]string

//Add adds an error message to a form field
func (err errors) Add(field, message string) {
	err[field] = append(err[field], message)
}

//Get returns the first error message
func (err errors) Get(field string) string {
	es := err[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
