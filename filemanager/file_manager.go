package filemanager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go_scenario_improver/model"

	"github.com/phpdave11/gofpdf"
)

// ListFiles returns all files (not directories) in the given directory.
func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read directory: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

// fixGarbledSymbols replaces common mis-encoded sequences with correct characters.
func fixGarbledSymbols(s string) string {
	replacements := map[string]string{
		"â€“": "–",
		"â€¦": "…",
		"â€˜": "‘",
		"â€™": "’",
		"â€œ": "“",
		"â€": "”",
		"â€¢": "•",
	}

	for bad, good := range replacements {
		s = strings.ReplaceAll(s, bad, good)
	}
	return s
}

// ReadFilesConcurrently reads all files from the given directory concurrently,
// returning a slice of Scenario structs containing filename and content.
func ReadFilesConcurrently(dir string, fileNames []string) ([]model.Scenario, error) {
	var wg sync.WaitGroup
	scenarioCh := make(chan model.Scenario, len(fileNames))
	errCh := make(chan error, len(fileNames))

	for _, name := range fileNames {
		wg.Add(1)
		go func(fname string) {
			defer wg.Done()
			contentRaw, err := os.ReadFile(filepath.Join(dir, fname))
			if err != nil {
				errCh <- fmt.Errorf("read file %s: %w", fname, err)
				return
			}
			content := fixGarbledSymbols(string(contentRaw))
			scenarioCh <- model.Scenario{Filename: fname, Content: content}
		}(name)
	}

	wg.Wait()
	close(scenarioCh)
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
		fmt.Fprintln(os.Stderr, err)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("one or more files failed to read")
	}

	var scenarios []model.Scenario
	for s := range scenarioCh {
		scenarios = append(scenarios, s)
	}

	return scenarios, nil
}

// SaveScenarioPDF saves the improved script content as a PDF file in outputDir.
// The output filename is derived from inputFilename by replacing the extension with ".pdf".
func SaveScenarioPDF(outputDir, inputFilename, content string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Improved Script - "+inputFilename, false)
	pdf.SetAuthor("YouTube Script AI", false)

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	title := fmt.Sprintf("Improved Script: %s", inputFilename)
	pdf.Cell(0, 10, title)
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)

	paragraphs := strings.Split(content, "\n")

	for _, para := range paragraphs {
		para = strings.TrimSpace(para)
		if para == "" {
			pdf.Ln(6)
			continue
		}

		if strings.HasPrefix(para, "[") && strings.HasSuffix(para, "]") {
			pdf.SetTextColor(220, 50, 50)
			pdf.SetFont("Arial", "B", 12)
			pdf.CellFormat(0, 8, para, "", 1, "", false, 0, "")
			pdf.SetTextColor(0, 0, 0)
			pdf.SetFont("Courier", "", 12)
		} else {
			pdf.MultiCell(0, 7, para, "", "L", false)
		}
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	outputPath := filepath.Join(outputDir, strings.TrimSuffix(inputFilename, filepath.Ext(inputFilename))+".pdf")
	return pdf.OutputFileAndClose(outputPath)
}
