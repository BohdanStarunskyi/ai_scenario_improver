# Go Scenario Improver

A concurrent Go application that automatically improves YouTube scripts using AI. Reads text files from an input directory, sends them to Google's Gemini AI for enhancement, and generates professional PDF outputs with improved grammar, style, and YouTube-style delivery notations.

## Project Description

* Processes multiple script files concurrently for maximum efficiency
* Automatically improves grammar, style, and flow to sound like a charismatic YouTuber
* Adds professional delivery notations like [Pause], [Emotional], [Excited], [Whisper], or [Joke]
* Generates beautifully formatted PDF files with proper styling and structure
* Uses Google's Gemini 2.0 Flash AI model for high-quality script improvements

## Features

* Concurrent file processing with goroutines and `sync.WaitGroup`
* AI-powered script enhancement using Google Gemini API
* Automatic PDF generation with professional formatting
* Support for multiple input files in batch processing
* Error handling and graceful failure recovery
* Environment-based API key configuration

## Tech Stack

* **Go** (Golang) 1.24.0
* **Google Gemini AI** — [https://generativelanguage.googleapis.com/](https://generativelanguage.googleapis.com/) for script improvement
* **gofpdf** — [https://github.com/phpdave11/gofpdf](https://github.com/phpdave11/gofpdf) for PDF generation
* **godotenv** — [https://github.com/joho/godotenv](https://github.com/joho/godotenv) for environment variable management
* Standard Go libraries for concurrency, I/O, and HTTP requests

## Project Structure

```
go_scenario_improver/
├── input/              # Input text files (scripts to improve)
├── result/             # Generated PDF files
├── filemanager/        # File I/O operations
├── model/              # Data structures
├── netowrking/         # AI API communication
├── main.go             # Entry point and orchestration
├── go.mod / go.sum     # Dependency management
└── .env                # Environment variables (API_KEY)
```

## How it Works (high level)

1. The program scans the `input/` directory for all text files
2. For each file, it spawns a goroutine that:
   * Reads the file content
   * Sends it to Google Gemini AI with specific instructions for YouTube script improvement
   * Receives the enhanced script with delivery notations
   * Generates a professional PDF with proper formatting
3. All files are processed concurrently for maximum speed
4. Results are saved as PDF files in the `result/` directory
5. Error handling ensures partial failures don't stop the entire process

## Example Input File

**File:** `input/cool_scenario.txt`

```
Hey! Today I want to show you three tech gadgets that might just blow your mind. First, there's this tiny smart ring that lets you control your phone with just a tap—whether it's changing songs, answering calls, or even unlocking your door without pulling out your phone. Then, imagine a full keyboard that folds up small enough to fit in your pocket. You can set it up anywhere and instantly turn any surface into your workspace. Finally, check out this futuristic 3D motion mouse. Instead of moving it on a desk, you move your hand in the air to control your cursor—like something straight out of a sci-fi movie. Which one would you want to try? Let me know down below!
```

## Example Output PDF



## How to Run

1. **Install Go** — [https://golang.org/dl/](https://golang.org/dl/)
2. **Clone the repository**

```bash
git clone https://github.com/your-username/go-scenario-improver.git
cd go-scenario-improver
```

3. **Set up your API key**
   * Create a `.env` file in the project root
   * Add your Google Gemini API key: `API_KEY=your_api_key_here`
   * Get your API key from [Google AI Studio](https://makersuite.google.com/app/apikey)

4. **Add your script files** — place `.txt` files in the `input/` directory

5. **Download dependencies**

```bash
go mod tidy
```

6. **Run**

```bash
go run main.go
```

7. **Check results** — find your improved PDFs in the `result/` directory

## API Key Setup

1. Visit [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Create a new API key
3. Create a `.env` file in your project root
4. Add: `API_KEY=your_actual_api_key_here`

## Configuration

The application uses these constants (defined in `main.go`):
* `inputDir = "./input"` — Directory containing script files
* `outputDir = "./result"` — Directory for generated PDFs

## Error Handling

* Individual file failures don't stop processing of other files
* Detailed error messages for debugging
* Graceful handling of API rate limits and network issues
* Automatic directory creation for output files

## Notes & Tips

* The AI model is specifically instructed to improve scripts for YouTube content
* Delivery notations are automatically added where appropriate
* PDFs are formatted for easy reading during video recording
* Consider the API rate limits when processing many files
* The application processes files concurrently, so large batches will be faster

## License

MIT — feel free to reuse and modify.
