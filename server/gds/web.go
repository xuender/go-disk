package gds

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/syndtr/goleveldb/leveldb"
)

// Web 服务
type Web struct {
	db *leveldb.DB // 数据库
}

// Close 关闭
func (w *Web) Close() error {
	log.Println("关闭Web服务")
	return w.db.Close()
}

// Start starts an HTTP server.
func (w *Web) Start(address string) error {
	return w.initEcho().Start(address)
}

func (w *Web) initEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = w.httpErrorHandler
	// 开发模式
	// if w.Dev {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	middleware.DefaultLoggerConfig.Format = `${time_rfc3339_nano} [${remote_ip}] ${host}(${method})${uri}(${status}) ${error} ${latency} ` +
		`[${latency_human}] IN:${bytes_in} OUT:${bytes_out}` + "\n"
	e.Use(middleware.Recover())
	// 跨域访问
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
		AllowCredentials: true,
	}))
	// } else {
	// 	e.HidePort = true
	// 	if f, err := os.OpenFile(w.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
	// 		middleware.DefaultLoggerConfig.Output = f
	// 		log.SetOutput(f)
	// 	}
	// }
	e.Use(middleware.Logger())
	// 二维码访问
	// e.GET("/qr", func(c echo.Context) error {
	// 	code, err := qr.Encode(w.URL, qr.Q)
	// 	if err != nil {
	// 		return c.String(http.StatusInternalServerError, "QR码生成错误: "+err.Error())
	// 	}
	// 	return c.Blob(http.StatusOK, "image/png", code.PNG())
	// })
	// 应用信息
	// e.GET("/about", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, w.App)
	// })
	// e.GET("/login", w.login) // 登录
	api := e.Group("/api") // API
	// // 需要身份认证
	// api.Use(middlewareJWT(w, "HS256"))
	w.filesRoute(api.Group("/files")) // 文件

	// // 静态资源
	// if w.Dev {
	// 	e.Static("/", "www")
	// } else {
	// 	e.Use(static.ServeRoot("/", getAssets("www")))
	// }
	log.Println("Go Disk 启动...")
	return e
}
func (w *Web) httpErrorHandler(err error, c echo.Context) {
	var code = http.StatusInternalServerError
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			if err := c.NoContent(code); err != nil {
				c.Logger().Error(err)
			}
		} else {
			if es, ok := err.(*echo.HTTPError); ok {
				if es.Code == 404 {
					c.Redirect(http.StatusMovedPermanently, "/")
					return
				}
				if err := c.JSON(es.Code, newHTTPError(es)); err != nil {
					c.Logger().Error(err)
				}
			} else {
				if err := c.JSON(code, newHTTPError(err)); err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}
}

// NewWeb 新建Web服务
func NewWeb(db string) (*Web, error) {
	web := &Web{}
	var err error
	if web.db, err = leveldb.OpenFile(db, nil); err != nil {
		return nil, err
	}
	return web, nil
}
