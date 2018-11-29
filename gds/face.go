package gds

import (
	"github.com/Kagami/go-face"
)

// Face 脸
type Face struct {
	face.Face `json:"-"` // 脸
	PeopleID  []byte     `json:"peopleID"` // 人脸
	FileID    []byte     `json:"fileID"`   // 文件
}
