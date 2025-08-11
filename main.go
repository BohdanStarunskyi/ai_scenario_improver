package main

import (
	"fmt"
	"go_scenario_improver/filemanager"
	"go_scenario_improver/model"
	"go_scenario_improver/netowrking"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const (
	inputDir  = "./input"
	outputDir = "./result"
)

func main() {
	godotenv.Load()
	if err := run(inputDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Files were saved to %v directory!", outputDir)
}

func run(dir string) error {
	fileNames, err := filemanager.ListFiles(dir)
	if err != nil {
		return err
	}

	scenarios, err := filemanager.ReadFilesConcurrently(dir, fileNames)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(scenarios))

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("api key wasn't set")
	}
	client := netowrking.NetworkManager{
		ApiKey: apiKey,
		Client: &http.Client{},
	}

	for _, sc := range scenarios {
		wg.Add(1)
		go func(s model.Scenario) {
			defer wg.Done()
			respContent, err := client.SendRequest(s.Content)
			if err != nil {
				errCh <- fmt.Errorf("request for %s: %w", s.Filename, err)
				return
			}
			if err := filemanager.SaveScenarioPDF(outputDir, s.Filename, respContent); err != nil {
				errCh <- fmt.Errorf("save PDF for %s: %w", s.Filename, err)
				return
			}
		}(sc)
	}

	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		for e := range errCh {
			fmt.Fprintln(os.Stderr, e)
		}
		return fmt.Errorf("one or more errors occurred during processing")
	}

	return nil
}
