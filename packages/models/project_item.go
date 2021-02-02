package models

import (
	"fmt"
	"github.com/strongo/validation"
	"strings"
)

// ProjectItem base class with ID and Title
type ProjectItem struct {
	ID    string `json:"id,omitempty" firestore:"id,omitempty" yaml:"id"`
	Title string `json:"title,omitempty" firestore:"title,omitempty" yaml:"title,omitempty"`
}

// MaxTitleLength defines maximum length of a title = 100
const MaxTitleLength = 100

func validateStringField(name, value string, isRequired bool, maxLen int) error {
	if isRequired && strings.TrimSpace(value) == "" {
		return validation.NewErrRecordIsMissingRequiredField(name)
	}
	if maxLen > 0 {
		if l := len(value); l > maxLen {
			return validation.NewErrBadRecordFieldValue(name,
				fmt.Sprintf("exceeds max length (%v): %v", maxLen, l))
		}
	}
	return nil
}

// Validate return error if not valid
func (v ProjectItem) Validate(isTitleRequired bool) error {
	if v.ID == "" {
		return validation.NewErrRecordIsMissingRequiredField("id")
	}
	if err := validateStringField("title", v.Title, isTitleRequired, MaxTitleLength); err != nil {
		return err
	}
	return nil
}