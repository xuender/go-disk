package gds

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/go-kit"
	"github.com/xuender/go-utils"
)

// Files 目录
type Files struct {
	DB        *kit.DB // 数据库
	TempPath  string  // 临时目录
	FilesPath string  // 文件目录
}

// List 文件列表
func (d *Files) List(dir string) (files []File) {
	files = []File{}
	d.DB.Get(utils.PrefixBytes([]byte(dir), _DirPrefix), &files)
	log.Println("getFiles:", len(files))
	return
}

// AddFile 增加文件
func (d *Files) AddFile(dir string, file File) error {
	ids := d.List(dir)
	ids = append(ids, file)
	return d.DB.Put(utils.PrefixBytes([]byte(dir), _DirPrefix), ids)
}

// TempFile 临时文件
func (d *Files) TempFile() (path string) {
	id := utils.NewID(_FilePrefix)
	path = filepath.Join(d.TempPath, id.String())
	return
}

// Save 保存文件
func (d *Files) Save(file, name, dir string, mod, size int64) error {
	// 保存文件,记录文件ID
	fid, err := utils.NewFileID(file)
	if err != nil {
		return err
	}
	fidBs := utils.PrefixBytes(fid.ID(), _FilePrefix)
	var data *File
	if has, err := d.DB.Has(fidBs); err == nil && !has {
		// 创建File
		if data, err = NewFile(file, name, size); err != nil {
			return err
		}
		data.ID = fid.ID()
		data.Mod = time.Unix(mod, 0)

		// 重命名
		path, f := d.getName(fid)
		kit.Mkdir(path)
		end := filepath.Join(path, f)
		os.Rename(file, end)
		d.DB.Put(fidBs, data)
		if data.Sub == JPEG {
			// TODO 人脸识别
			go func() {
				if faces, err := _rec.RecognizeFile(end); err == nil {
					d.DB.Put(utils.PrefixBytes(fid.ID(), _RecognitionPrefix), faces)
				}
			}()
		}
	} else {
		// 删除
		os.Remove(file)
		data = &File{}
		d.DB.Get(fidBs, data)
		data.Name = name
		// TODO 重名检查,名称修改
	}
	d.AddFile(dir, *data)
	return nil
}

func (d *Files) getName(id *utils.FileID) (dir, name string) {
	str := id.String()
	dir = filepath.Join(d.FilesPath, str[:3], str[3:6])
	name = str[6:]
	return
}

func init() {
	PutRoute("files", func(g *echo.Group) {
		// 文件列表
		g.GET("/:dir", func(c echo.Context) error {
			dir := c.Param("dir")
			log.Println("读取目录", dir)
			files := _files.List(dir)
			return c.JSON(http.StatusOK, files)
		})
		// 上传文件
		g.POST("/:dir", func(c echo.Context) error {
			dir := c.Param("dir")
			log.Println("上传目录:", dir)
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
			mod, _ := strconv.ParseInt(c.FormValue("mod"), 10, 64)
			if err := _files.Save(f, file.Filename, dir, mod, file.Size); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, "ok")
		})
	})
}
