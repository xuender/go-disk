package gds

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
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
	Size int       `json:"size,omitempty"` // 尺寸
	Ca   time.Time `json:"ca"`             // 创建时间
	Mod  time.Time `json:"mod"`            // 修改时间
}

// 文件路由
func (w *Web) filesRoute(c *echo.Group) {
	// c.GET("", w.customersGet)                                                     // 客户列表
	// c.POST("", w.customerPost)                                                    // 客户创建
	// c.PUT("/:id", w.customerPut)                                                  // 客户修改
	// c.DELETE("/:id", w.customerDelete)                                            // 删除客户
	// c.DELETE("", w.customersDelete)                                               // 清除客户
	// c.POST("/file", w.customersFile)                                              // 上传客户文件
	c.GET("/:file", func(c echo.Context) error {
		fs := []File{
			File{Name: "测试.doc", Size: 50000, Type: FILE, Ca: time.Now(), Mod: time.Now()},
			File{Name: "目录", Type: DIR, Ca: time.Now(), Mod: time.Now()},
		}
		return c.JSON(http.StatusOK, fs)
	})
}
