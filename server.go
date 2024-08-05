// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
//
// See golang.org/x/example/outyet for a more sophisticated server.
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"langchain-echo-deepinfra/modules/coreSetup"
	"langchain-echo-deepinfra/modules/llms"
)

func main() {

	coreSetup.LoadEnvVariable("dev")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Outer World!")
	})

	// Register the GenerateCompletion handler
	e.POST("/generateCompletions", llms.GenerateCompletion)

	// Register the GenerateCompletion handler
	e.GET("/generateEmbeddings", llms.GenerateEmbedding)

	e.Logger.Fatal(e.Start(":3000"))

}
