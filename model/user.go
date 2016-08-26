package model

import "github.com/xeipuuv/gojsonschema"

type (
	User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

// Valid return User model invalid error
func (user *User) Valid() (bool, []string) {
	schema := gojsonschema.NewReferenceLoader("file:///Users/himu62/go/src/github.com/himu62/echo-testing/jsonschema/user.json")
	data := gojsonschema.NewGoLoader(user)

	result, err := gojsonschema.Validate(schema, data)
	if err != nil {
		return false, []string{err.Error()}
	}

	if !result.Valid() {
		errs := make([]string, 0, 5)
		for _, desc := range result.Errors() {
			errs = append(errs, desc.Description())
		}
		return false, errs
	}
	return true, nil
}
