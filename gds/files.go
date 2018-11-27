package gds

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
	"github.com/xuender/go-kit"
	"github.com/xuender/go-utils"
)

// Files 目录
type Files struct {
	TempPath  string // 临时目录
	FilesPath string // 文件目录
}

// List 文件列表
func (d *Files) List(dir string) (files []File) {
	key := utils.PrefixBytes([]byte(dir), _DirPrefix)
	log.Printf("%s, %x\n", dir, key)
	files = []File{}
	if has, err := _db.Has(key); err == nil && has {
		ids := [][]byte{}
		_db.Get(key, &ids)
		for _, id := range ids {
			f := File{}
			if _db.Get(id, &f) == nil {
				files = append(files, f)
			}
		}
	} else {
		m := map[string]bool{}
		_db.Iterator(key, func(k, value []byte) bool {
			log.Printf("%v\n", k)
			m[subName(string(k), len(key))] = true
			return false
		})
		for name := range m {
			d := File{
				Name: name,
				Type: DIR,
			}
			files = append(files, d)
		}
	}
	return
}
func subName(name string, size int) string {
	s := name[size:]
	if s[0] == '/' {
		s = s[1:]
	}
	i := strings.Index(s, "/")
	if i < 0 {
		return s
	}
	return s[:i]
}

// AddFile 增加文件
func (d *Files) AddFile(dir string, fid []byte) error {
	key := utils.PrefixBytes([]byte(dir), _DirPrefix)
	ids := [][]byte{}
	_db.Get(key, &ids)
	ids = append(ids, fid)
	return _db.Put(key, ids)
}

// TempFile 临时文件
func (d *Files) TempFile() (path string) {
	id := utils.NewID(_FilePrefix)
	path = filepath.Join(d.TempPath, id.String())
	return
}

// Save 保存文件
func (d *Files) Save(file, name string) error {
	// 保存文件,记录文件ID
	fid, err := utils.NewFileID(file)
	if err != nil {
		return err
	}
	fidBs := utils.PrefixBytes(fid.ID(), _FilePrefix)
	var data *File
	// 不存在
	if has, err := _db.Has(fidBs); err == nil && !has {
		// 创建File
		if data, err = NewFile(file, name); err != nil {
			return err
		}
		data.ID = fidBs
		// 重命名
		kit.Mkdir(fmt.Sprintf("%s/%s", d.FilesPath, data.Dir()))
		os.Rename(file, fmt.Sprintf("%s/%s/%s", d.FilesPath, data.Dir(), data.FileName()))
		_db.Put(fidBs, data)
		d.AddFile(data.Dir(), fidBs)
	} else {
		// 删除
		os.Remove(file)
	}
	return nil
}

func init() {
	PutRoute("files", func(g *echo.Group) {
		// 文件列表
		g.GET("", func(c echo.Context) error {
			dir := c.QueryParam("dir")
			log.Println("读取目录", dir)
			files := _files.List(dir)
			return c.JSON(http.StatusOK, files)
		})
		// 上传文件
		g.POST("", func(c echo.Context) error {
			log.Println("上传")
			// 来源
			file, err := c.FormFile("uploadfile")
			if err != nil {
				return err
			}
			// log.Println(file.Header)
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			// 目的
			f := _files.TempFile()
			dst, err := os.Create(f)
			if err != nil {
				return err
			}
			defer dst.Close()
			// 复制
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
			if err := _files.Save(f, file.Filename); err != nil {
				return err
			}
			// TODO 文件尺寸校验,防止错误
			return c.JSON(http.StatusOK, "ok")
		})
	})
}
