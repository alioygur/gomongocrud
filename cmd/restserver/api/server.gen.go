// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /tasks)
	FindTasks(ctx echo.Context) error

	// (POST /tasks)
	AddTask(ctx echo.Context) error

	// (DELETE /tasks/{id})
	DeleteTask(ctx echo.Context, id string) error

	// (GET /tasks/{id})
	FindTaskById(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindTasks converts echo context to params.
func (w *ServerInterfaceWrapper) FindTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindTasks(ctx)
	return err
}

// AddTask converts echo context to params.
func (w *ServerInterfaceWrapper) AddTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddTask(ctx)
	return err
}

// DeleteTask converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTask(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTask(ctx, id)
	return err
}

// FindTaskById converts echo context to params.
func (w *ServerInterfaceWrapper) FindTaskById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindTaskById(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/tasks", wrapper.FindTasks)
	router.POST("/tasks", wrapper.AddTask)
	router.DELETE("/tasks/:id", wrapper.DeleteTask)
	router.GET("/tasks/:id", wrapper.FindTaskById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RVwW7jNhD9FWLao2q5TU+6JXFb+NIUaXoKjGIijmwmEsnljJIYhv59QdLaxJGR3WB3",
	"gQB7Mk2N3rx5b2a0g9p13lmywlDtgOsNdZiOf4TgQjz44DwFMZSua6cp/srWE1TAEoxdw1BAR8y4PvZs",
	"KCDQh94E0lBdZ4Sn+FUxxrubW6olYv1ND1fId9PsmrgOxotx9igJi90XMEhRxQHYMRYjBWzbiwaq6x38",
	"HKiBCn4qn1Qr95KVI+ehmEgWCIX0/yjxX+NCF0+gUegXMYnKpBCjj9bXe/1GqBelGw3Fc0IHkFMRVkME",
	"MLZxE/XhVDF2viV1+s9SyQZF9UysUHkSFhdIISu0ih5zmDilqXOWJaCQagilD8TKWCUbUheebEQ6mc0V",
	"e6pNY2pMqQoQI22k9W9G+su1aNfq/PK/RUwOBdxT4Mzq19l8No9SOU8WvYEKTtJVAR5lkwwpBfkundYk",
	"08IuSfpgWWHbqhyZ0EJis9RQwZ/G6qv9k0DsneVs9W/zeR4SK2QTNHrf7gspbzm3bW6ZeDJCXXrxtc7a",
	"t9XoDYaAW0jGHPKOXNVIB9LjBvtW3sToNSJ5JRzJ3Ft69FQLaUVPMd7xEXXPU/OxsvSQ1J2Ie6qTtpAb",
	"l1jOnN5+sxo+jem0ingfuxS1hudTI6Gn4SuN/ry/793PodjPTbkzesi+tiQ0dTjfx03Axq7j5MdCbpBJ",
	"K5eHfblQ3EfepCf+L9Lr+xbwGLAjocBpAR8mWi6UazJ6Wi6JTlxXUKVhh/GDkNfeoaPFM5Fe7szVxO3f",
	"p2WmvDmpfgfT9uoqc5aOD9u4yc62S/02uRuSevOd1P7hZmsogCncj7r3oYUKNiK+KsvxczrjB1yvKcyM",
	"K+OHbVgNHwMAAP//qlnMxL8JAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
