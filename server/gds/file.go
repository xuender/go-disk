package gds

import (
	"time"
)

const (
	// DIR 目录
	DIR = iota
	// FILE 文件
	FILE
)

// File 文件
type File struct {
	Name string    `json:"name"`           // 名称
	Type int       `json:"type"`           // 类型
	Size int64     `json:"size,omitempty"` // 尺寸
	Ca   time.Time `json:"ca"`             // 创建时间
	Mod  time.Time `json:"mod"`            // 修改时间
	ID   []byte    `json:"id"`             // 文件ID
}
