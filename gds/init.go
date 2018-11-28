package gds

import (
	"log"
	"path/filepath"

	"github.com/xuender/go-utils"

	"github.com/Kagami/go-face"
	"github.com/xuender/go-kit"
)

const (
	_FilePrefix   = 'F' // 文件前缀
	_DirPrefix    = 'D' // 目录前缀
	_PeoplePrefix = 'P' // 人
)

var _files *Files         // 文件管理
var _rec *face.Recognizer // 人脸识别
var _db *kit.DB           // 数据库
var _peoples Peoples      // 人脸
// InitDB 数据库初始化
func InitDB(db *kit.DB) {
	_db = db
}

// Init 初始化
func Init(dataDir string) error {
	tempPath := filepath.Join(dataDir, "temp")
	photoPath := filepath.Join(dataDir, "photo")
	log.Println(tempPath, photoPath)
	if err := kit.Mkdir(tempPath); err != nil {
		return err
	}
	_files = &Files{
		TempPath:  tempPath,
		PhotoPath: photoPath,
	}
	return nil
}

// InitRec 初始化人脸识别
func InitRec(rec *face.Recognizer) {
	_rec = rec
	_peoples = []People{}
	_db.Iterator([]byte{_PeoplePrefix, '-'}, func(key, value []byte) bool {
		p := People{}
		if utils.Decode(value, &p) == nil {
			_peoples = append(_peoples, p)
		}
		return false
	})
	if len(_peoples) > 0 {
		_rec.SetSamples(_peoples.Samples(), _peoples.Cats())
	}
}
