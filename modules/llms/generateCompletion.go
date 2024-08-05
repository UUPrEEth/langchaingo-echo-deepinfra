// This method is used to generate LLM completions
package llms

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tmc/langchaingo/llms"
)

// Creating the Request Struct
var requestData struct {
	Prompt string `json:"prompt"`
}

func GenerateCompletion(e echo.Context) error {

	if err := json.NewDecoder(e.Request().Body).Decode(&requestData); err != nil {
		return e.JSON(400, echo.Map{"error": "Invalid JSON"})
	}

	ctx := context.Background()
	// llm, err := openai.New()

	// if err != nil {
	// 	log.Fatal(err)
	// 	return e.JSON(500, echo.Map{"error": err.Error()})
	// }

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

// Generic Usecase
// jsonBody := make(map[string]interface{})
// if err := json.NewDecoder(e.Request().Body).Decode(&jsonBody); err != nil {
// 	e.JSON(400, echo.Map{"error": "Invalid JSON"})
// 	return
// }
