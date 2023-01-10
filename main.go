package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/hello", Hello())
    e.GET("/api/hello", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "AOI-1up: Hello, CI/CD! ver.3")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "AOI-1up", "message": "Hello, CI/CD! ver.3"})
    }
}