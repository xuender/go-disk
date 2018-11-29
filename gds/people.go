package gds

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	face "github.com/Kagami/go-face"
	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// People 人
type People struct {
	Name  string `json:"name"`  // 姓名
	Faces []Face `json:"faces"` // 脸
}

// GetFace 获取面孔
func (p *People) GetFace(i int) ([]byte, error) {
	f := p.Faces[i]
	file := File{}
	if _db.Get(f.FileID, &file) == nil {
		return SubImage(filepath.Join(_files.PhotoPath, file.Dir(), file.FileName()), f.Rectangle)
	}
	return nil, nil
}

// SubImage 子图像
func SubImage(photo string, rectangle image.Rectangle) (bs []byte, err error) {
	reader, err := os.Open(photo)
	if err != nil {
		return
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		return
	}
	rgbImg := m.(*image.YCbCr)
	buff := bytes.NewBuffer(nil)
	jpeg.Encode(buff, rgbImg.SubImage(rectangle), nil)
	bs = buff.Bytes()
	return
}

// Peoples 人脸识别记录
type Peoples []People

// IsNew 新人脸识别记录
func (ps Peoples) IsNew() bool {
	return len(ps) == 0
}

// Get 获取ID
func (ps Peoples) Get(cid int) []byte {
	i := 0
	for _, p := range ps {
		for _, f := range p.Faces {
			if cid == i {
				return f.PeopleID
			}
			i++
		}
	}
	return nil
}

// Cats 标签
func (ps Peoples) Cats() []int32 {
	i := 0
	cats := []int32{}
	for _, p := range ps {
		for f := 0; f < len(p.Faces); f++ {
			cats = append(cats, int32(i))
			i++
		}
	}
	return cats
}

// Samples 例子
func (ps Peoples) Samples() []face.Descriptor {
	samples := []face.Descriptor{}
	for _, p := range ps {
		for _, f := range p.Faces {
			samples = append(samples, f.Descriptor)
		}
	}
	return samples
}

// PeopleID 获取人脸标识
func (ps Peoples) PeopleID(face face.Face, fid []byte) []byte {
	log.Println("IsNew", ps.IsNew())
	// if !ps.IsNew() {
	// 	cid := _rec.Classify(face.Descriptor)
	// 	log.Println("IsNew", ps.IsNew(), cid)
	// 	if cid >= 0 {
	// 		return _peoples.Get(cid)
	// 	}
	// }
	id := utils.NewID(_PeoplePrefix)
	p := People{
		Name:  "未知",
		Faces: []Face{Face{Face: face, PeopleID: id[:], FileID: fid}},
	}
	log.Println("新增 People", len(ps))
	_db.Put(id[:], p)
	ps = append(ps, p)
	log.Println("新增 People", len(ps))
	_rec.SetSamples(ps.Samples(), ps.Cats())
	return nil
}

func init() {
	PutRoute("peoples", func(g *echo.Group) {
		// 人脸列表
		g.GET("", func(c echo.Context) error {
			log.Println("人脸列表")
			return c.JSON(http.StatusOK, _peoples)
		})
		// 显示面孔
		g.GET("/:index", func(c echo.Context) error {
			idStr := c.QueryParam("id")
			log.Println("id", idStr)
			id, err := base64.StdEncoding.DecodeString(idStr)
			if err != nil {
				return err
			}
			index := c.Param("index")
			log.Println("id", id, index)
			p := People{}
			if _db.Get(id, &p) == nil {
				i, err := strconv.Atoi(index)
				if err != nil {
					return err
				}
				bs, err := p.GetFace(i)
				if err != nil {
					return err
				}
				return c.Blob(http.StatusOK, "image/jpeg", bs)
			}
			return c.JSON(http.StatusOK, "xx")
		})
	})
}
