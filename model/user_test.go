package model

import (
	"reflect"
	"testing"
)

func TestValid(t *testing.T) {
	seeds := []User{
		User{
			Name:  "hogehoge",
			Email: "hogehoge@hoge.com",
		},
	}

	type Correct struct {
		Valid  bool
		Errors []string
	}
	corrects := []Correct{
		Correct{
			Valid: true,
			// Errors is nil
		},
	}

	for i, v := range seeds {
		valid, errs := v.Valid()

		if valid != corrects[i].Valid {
			t.Errorf("valid should be %t, but got %t", corrects[i].Valid, valid)
		}
		if !reflect.DeepEqual(errs, corrects[i].Errors) {
			t.Errorf("errors should be %v, but got %v", corrects[i].Errors, errs)
		}
	}
}
