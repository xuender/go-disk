package gds

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	_, err := NewFile("../../LICENSE", "LICENSE")
	assert.NotNil(t, err)
}

func ExampleNewFile() {
	ca, _ := time.Parse("2006-01-02", "2018-01-15")
	f := File{
		Ca: ca,
	}
	fmt.Println(f.Dir())

	// Output:
	// 2018/01/15
}
