package main

import (
	"github.com/labstack/echo/v4"
	"vivasoft-employee-entry-time-management/package/containers"
)

func main() {
	e := echo.New()
	containers.Serve(e)
}
