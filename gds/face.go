package gds

import (
	"github.com/Kagami/go-face"
)

// Face 脸
type Face struct {
	face.Face        // 脸
	PeopleID        []byte // 人脸
}
