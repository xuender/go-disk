package gds

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// WebStart starts an HTTP server.
func WebStart(address string) error {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = httpErrorHandler
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	middleware.DefaultLoggerConfig.Format = `${time_rfc3339_nano} [${remote_ip}] ${host}(${method})${uri}(${status}) ${error} ${latency} ` +
		`[${latency_human}] IN:${bytes_in} OUT:${bytes_out}` + "\n"
	e.Use(middleware.Recover())
	// 支持跨域访问
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
		AllowCredentials: true,
	}))
	// 日志
	e.Use(middleware.Logger())
	api := e.Group("/api")
	// TODO 需要身份认证

	for k, f := range _routes {
		if !strings.HasPrefix(k, "/") {
			k = "/" + k
		}
		f(api.Group(k))
	}
	// 静态资源处理
	e.Static("/", "www")
	log.Println("Go Disk 启动...")
	return e.Start(address)
}

// 路由表
var _routes = map[string]func(*echo.Group){}

// PutRoute 设置路由
func PutRoute(path string, f func(*echo.Group)) error {
	if _, has := _routes[path]; has {
		return fmt.Errorf("路由错误 %s", path)
	}
	_routes[path] = f
	return nil
}

func httpErrorHandler(err error, c echo.Context) {
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
