package gds

import (
	"log"

	"github.com/xuender/go-kit"
)

const (
	_FilePrefix = 'F' // 文件前缀
	_DirPrefix  = 'D' // 目录前缀
)

var _files *Files

// Init 初始化
func Init(db *kit.DB, tempPath, filesPath string) error {
	log.Println(tempPath, filesPath)
	if err := kit.Mkdir(tempPath); err != nil {
		return err
	}
	_files = &Files{
		DB:        db,
		TempPath:  tempPath,
		FilesPath: filesPath,
	}
	return nil
}