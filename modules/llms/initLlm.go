// This method is used to initiate LLM with openai
package llms

import (
	"log"

	"github.com/tmc/langchaingo/llms/openai"
)

// in this methodology, the same llm is used in multiple functions
var llm *openai.LLM

func init() {
	// Initialize LLM here
	var err error
	llm, err = openai.New()
	if err != nil {
		log.Fatal(err)
	}
}
