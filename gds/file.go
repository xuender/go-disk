package gds

import (
	"fmt"
	"time"

	"github.com/xuender/go-kit"
)

const (
	// DIR 目录
	DIR = iota
	// FILE 文件
	FILE
	// IMAGE 图像
	IMAGE
	// JPEG 照片
	JPEG
)

// File 文件
type File struct {
	Name  string    `json:"name"`           // 名称,可用于排序
	Type  int       `json:"type"`           // 类型
	Sub   int       `json:"sub,omitempty"`  // 子类型
	Size  int64     `json:"size,omitempty"` // 尺寸
	Ca    time.Time `json:"ca"`             // 创建时间
	ID    []byte    `json:"id"`             // 文件ID
	Faces []Face    `json:"faces"`          // 脸
}

// Dir 目录
func (f *File) Dir() string {
	return fmt.Sprintf("%d/%02d/%02d", f.Ca.Year(), f.Ca.Month(), f.Ca.Day())
}

// FileName 文件名
func (f *File) FileName() string {
	return fmt.Sprintf("%x.jpg", f.ID)
}

// NewFile 创建文件
func NewFile(path, name string, fid []byte) (file *File, err error) {
	photo, err := kit.NewPhoto(path, _rec)
	if err != nil {
		return
	}
	file = &File{
		Name: name,
		Type: IMAGE,
		Sub:  JPEG,
		Size: photo.Size,
		Ca:   photo.Ca,
	}
	file.Faces = []Face{}
	if photo.Faces != nil {
		for _, f := range photo.Faces {
			face := Face{}
			face.Face = f
			face.PeopleID = _peoples.PeopleID(f, fid)
			file.Faces = append(file.Faces, face)
		}
	}
	return
}
