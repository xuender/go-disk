package gds

import (
	"log"

	"github.com/Kagami/go-face"

	"github.com/xuender/go-kit"
)

const (
	_FilePrefix        = 'F' // 文件前缀
	_DirPrefix         = 'D' // 目录前缀
	_RecognitionPrefix = 'R' // 人脸识别数据
)

var _files *Files
var _rec *face.Recognizer // 人脸识别

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

// InitRec 初始化人脸识别
func InitRec(rec *face.Recognizer) {
	_rec = rec
}
