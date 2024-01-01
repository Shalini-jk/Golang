// example_test.go

package operation

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := -1

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
