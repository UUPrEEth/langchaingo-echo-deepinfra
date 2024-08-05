package llms

import (
	"langchain-echo-deepinfra/modules/coreSetup"
	"log"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/qdrant"
)

func GenerateEmbedding(e echo.Context) error {
	// llm, err := openai.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}

	qdrantURL, err := coreSetup.GetEnv("QDRANT_URL")
	if err != nil {
		log.Fatal(err)
	}

	qdrantAPIKey, err := coreSetup.GetEnv("QDRANT_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	url, err := url.Parse(qdrantURL)
	if err != nil {
		log.Fatal(err)
	}

	//store, err := qdrant.New(
	// since no new variable store is initialized, its = intead of :=
	_, err = qdrant.New(
		qdrant.WithURL(*url),
		qdrant.WithAPIKey(qdrantAPIKey),
		qdrant.WithCollectionName("notes"),
		qdrant.WithEmbedder(embedder),
	)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
