package gds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xuender/go-utils"
)

func TestFiles_getName(t *testing.T) {
	// 128583ac2d75ef253f5229b01ad994905d2c
	files := Files{
		FilesPath: "/tmp",
	}
	id, _ := utils.NewFileID("../../LICENSE")
	p, f := files.getName(id)
	assert.Equal(t, p, "/tmp/128/583")
	assert.Equal(t, f, "ac2d75ef253f5229b01ad994905d2c")
}

func ExampleFiles_getName() {
	id, _ := utils.NewFileID("../../LICENSE")
	fmt.Println(id)

	// Output:
	// 128583ac2d75ef253f5229b01ad994905d2c
}
