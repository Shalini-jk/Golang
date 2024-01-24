package main

import (
	"reflect"
	"testing"
)



func TestSkillList(t *testing.T) {
	expectedSkills := []string{"GoLang", "Python"}

	// Call the function
	resultSkills := skillList()

	// Compare the result with the expected value
	if !reflect.DeepEqual(resultSkills, expectedSkills) {
		t.Errorf("Expected skills %v, but got %v", expectedSkills, resultSkills)
	}
}
