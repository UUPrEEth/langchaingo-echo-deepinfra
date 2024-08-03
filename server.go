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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Outer World!")
	})

	e.POST("/generateCompletions", generateCompletion)

	e.Logger.Fatal(e.Start(":3000"))

}

func generateCompletion(e echo.Context) error {
	var requestData struct {
		Prompt string `json:"prompt"`
	}

	if err := json.NewDecoder(e.Request().Body).Decode(&requestData); err != nil {
		return e.JSON(400, echo.Map{"error": "Invalid JSON"})
	}

	// Generic Usecase
	// jsonBody := make(map[string]interface{})
	// if err := json.NewDecoder(e.Request().Body).Decode(&jsonBody); err != nil {
	// 	e.JSON(400, echo.Map{"error": "Invalid JSON"})
	// 	return
	// }

	ctx := context.Background()
	llm, err := openai.New()

	if err != nil {
		log.Fatal(err)
		return e.JSON(500, echo.Map{"error": err.Error()})
	}

	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, requestData.Prompt,
		llms.WithModel("mistralai/Mistral-7B-Instruct-v0.3"),
		llms.WithTemperature(0.2),
	)
	if err != nil {
		log.Fatal(err)
		return e.JSON(500, echo.Map{"error": err.Error()})
	}
	fmt.Println("completion succesful")

	return e.JSON(http.StatusOK, echo.Map{"completion": completion})

}
