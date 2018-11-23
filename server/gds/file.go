package gds

import (
	"fmt"
	"os"
	"time"

	filetype "gopkg.in/h2non/filetype.v1"
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
	Name string    `json:"name"`           // 名称
	Type int       `json:"type"`           // 类型
	Sub  int       `json:"sub,omitempty"`  // 子类型
	Size int64     `json:"size,omitempty"` // 尺寸
	Ca   time.Time `json:"ca"`             // 创建时间
	Mod  time.Time `json:"mod"`            // 修改时间
	ID   []byte    `json:"id"`             // 文件ID
}

// NewFile 创建文件
func NewFile(path, name string, size int64) (file *File, err error) {
	// size 尺寸校验
	s, err := os.Stat(path)
	if err != nil {
		return
	}
	file = &File{
		Name: name,
		Type: FILE,
		Size: size,
		Ca:   time.Now(),
	}
	if s.IsDir() {
		file.Size = 0
		file.Type = DIR
	} else {
		if s.Size() != size {
			err = fmt.Errorf("文件 %s 尺寸错误 %d != %d", path, s.Size(), size)
			return
		}
		// 文件类型判断
		t, err := filetype.MatchFile(path)
		if err == nil {
			if t.MIME.Type == "image" {
				file.Type = IMAGE
			}
			if t.MIME.Subtype == "jpeg" {
				file.Sub = JPEG
			}
		}
	}
	return
}
