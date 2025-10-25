# ScenarioAI - Script Improver

A full-stack web application that transforms your ideas into compelling scenarios using AI. Features a sleek React frontend and a robust Go backend API, powered by Google's Gemini 2.0 Flash for intelligent script enhancement with professional delivery notations.

## Project Description

* **Modern Web Interface**: Intuitive React frontend with real-time scenario generation
* **Powerful API Backend**: High-performance Go server with clean architecture and comprehensive testing
* **AI-Powered Enhancement**: Automatically improves grammar, makes it comfortable to make videos
* **Professional Delivery Notations**: Adds cues like [Pause], [Emotional], [Excited], [Whisper], or [Joke]
* **Real-time Processing**: Instant scenario generation with loading states and error handling
* **Responsive Design**: Beautiful UI that works seamlessly across all devices

## Features

### Frontend
* ✨ Beautiful, modern React interface with TypeScript
* 🎨 Gradient backgrounds with glass morphism design
* 📱 Fully responsive design for all screen sizes
* ⚡ Real-time API communication with loading states
* 🔄 Smooth animations and transitions
* 🚀 Built with Vite for optimal performance

### Backend
* 🏗️ Clean architecture with separation of concerns
* 🚀 High-performance Echo web framework
* 🔒 CORS support and rate limiting
* 📝 Comprehensive request validation
* 🧪 100% test coverage with unit tests
* 📊 Structured logging with detailed error tracking
* ⚡ Concurrent request handling

## Tech Stack

### Frontend
* **React** 19.2.0 with TypeScript
* **Vite** 7.1.12 for blazing fast development
* **CSS3** with modern features (gradients, backdrop-filter, animations)
* **ESLint** for code quality

### Backend
* **Go** 1.24.0 with Echo v4 web framework
* **Google Gemini 2.0 Flash AI** for script improvement
* **Structured Logging** with built-in slog
* **Environment Configuration** with godotenv
* **Request Validation** with go-playground/validator
* **Testing** with testify and mocks

## Project Structure

```
ai_scenario_improver/
├── backend/                    # Go API Server
│   ├── main.go                # Application entry point
│   ├── go.mod                 # Go module dependencies
│   ├── application/           # Application initialization
│   │   └── app.go            # Dependency injection container
│   ├── controllers/          # HTTP request handlers
│   │   ├── health.go         # Health check endpoints
│   │   ├── health_test.go    # Health controller tests
│   │   ├── scenario.go       # Scenario generation endpoints
│   │   └── scenario_test.go  # Scenario controller tests
│   ├── dto/                  # Data transfer objects
│   │   └── scenario.go       # Request/response structures
│   ├── model/               # Domain models
│   │   └── scenario.go      # Scenario data model
│   ├── services/            # Business logic layer
│   │   ├── scenario.go      # Scenario generation service
│   │   └── scenario_test.go # Service unit tests
│   ├── utils/               # Utility packages
│   │   └── network_manager.go # AI API communication
│   ├── mocks/               # Test mocks
│   │   └── mocks.go         # Service mocks for testing
│   ├── router/              # Route definitions
│   │   └── router.go        # API route setup
│   └── .env                 # Environment variables (API_KEY)
├── web/                     # React Frontend
│   ├── src/
│   │   ├── App.tsx          # Main application component
│   │   ├── App.css          # Application styles
│   │   ├── main.tsx         # React entry point
│   │   ├── index.css        # Global styles
│   │   └── assets/          # Static assets
│   ├── public/              # Public assets
│   ├── package.json         # Node.js dependencies
│   ├── tsconfig.json        # TypeScript configuration
│   ├── vite.config.ts       # Vite build configuration
│   └── index.html           # HTML template
├── LICENSE                  # MIT license
└── README.md               # This file
```

## How it Works

### Frontend Flow
1. **User Input**: User enters their video idea in the modern React interface
2. **Real-time Validation**: Input is validated on the client side
3. **API Request**: Frontend sends POST request to `/generate` endpoint
4. **Loading State**: Beautiful loading animation while processing
5. **Results Display**: Enhanced scenario appears with smooth animations
6. **Reset Option**: User can generate new scenarios with one click

### Backend Architecture
1. **Request Handling**: Echo router receives and validates requests
2. **Service Layer**: Business logic processes the scenario generation
3. **AI Integration**: Network manager communicates with Google Gemini API
4. **Response Processing**: AI response is validated and returned
5. **Error Handling**: Comprehensive logging and graceful error responses
6. **CORS & Security**: Built-in middleware for security and rate limiting

## Screenshots

### Frontend Interface
<img width="1440" height="812" alt="Screenshot 2025-10-25 at 10 30 25" src="https://github.com/user-attachments/assets/40eb8be6-7e7b-4872-b560-e2fb7145ac60" />
<img width="1440" height="818" alt="Screenshot 2025-10-25 at 10 30 38" src="https://github.com/user-attachments/assets/e6496885-4fac-4557-a591-6d227ee4c2e6" />

## Example Usage

**User Input:**
```
Hey! Today I want to show you three tech gadgets that might just blow your mind. First, there's this tiny smart ring that lets you control your phone with just a tap—whether it's changing songs, answering calls, or even unlocking your door without pulling out your phone. Then, imagine a full keyboard that folds up small enough to fit in your pocket. You can set it up anywhere and instantly turn any surface into your workspace. Finally, check out this futuristic 3D motion mouse. Instead of moving it on a desk, you move your hand in the air to control your cursor—like something straight out of a sci-fi movie. Which one would you want to try? Let me know down below!
```

**Generated Output:**
```
[Excited] Hey there, tech lovers! [Pause] Today I'm about to show you THREE absolutely mind-blowing gadgets that will make you question everything you know about technology! 

[Energetic] First up – this tiny smart ring that's basically magic on your finger! [Pause] We're talking about controlling your ENTIRE phone with just a simple tap. Change songs? Tap. Answer calls? Tap. [Whisper] Even unlock your front door without ever reaching for your phone. [Excited] It's like having superpowers!

[Pause] But wait, it gets better! [Emotional] Picture this: a full-sized keyboard that folds down so small, it literally fits in your pocket! [Pause] You can set up your workspace ANYWHERE – on a park bench, at a coffee shop, even on an airplane tray table. [Joke] Finally, a reason to be productive during those boring flights!

[Amazed] And here's where it gets sci-fi crazy... [Pause] This 3D motion mouse lets you control your cursor by moving your hand through the AIR! [Excited] No desk needed, no mouse pad – just pure hand gestures like you're casting spells! [Emotional] It's literally something straight out of Minority Report!

[Pause] So here's my question for YOU – [Engaged] which one of these futuristic gadgets would you want to get your hands on first? [Pause] Drop your thoughts in the comments below and let me know which one blew your mind the most! [Smile] And don't forget to hit that subscribe button for more incredible tech discoveries!
```



## Getting Started

### Prerequisites
* **Go** 1.24.0+ — [Download here](https://golang.org/dl/)
* **Node.js** 18+ — [Download here](https://nodejs.org/)
* **Google Gemini API Key** — [Get it here](https://makersuite.google.com/app/apikey)

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/BohdanStarunskyi/ai_scenario_improver.git
cd ai_scenario_improver
```

2. **Backend Setup**
```bash
# Navigate to backend directory
cd backend

# Install Go dependencies
go mod tidy

# Create .env file with your API key
echo "API_KEY=your_gemini_api_key_here" > .env

# Run tests to ensure everything works
go test ./...

# Start the backend server
go run main.go
```
The backend will start on `http://localhost:8080`

3. **Frontend Setup** (in a new terminal)
```bash
# Navigate to web directory
cd web

# Install Node.js dependencies
npm install

# Start the development server
npm run dev
```
The frontend will start on `http://localhost:5173`

### Running the Application

1. **Start Backend**: `cd backend && go run main.go` (Port 8080)
2. **Start Frontend**: `cd web && npm run dev` (Port 5173)
3. **Open Browser**: Navigate to `http://localhost:5173`
4. **Enter Your Idea**: Type your video concept and click "Generate Scenario"
5. **Get Enhanced Script**: Watch as AI transforms your idea into a professional script

## API Endpoints

### Backend API (Port 8080)

#### Health Check
```bash
GET /ping
Response: {"message": "pong"}
```

#### Generate Scenario
```bash
POST /generate
Content-Type: application/json

Request Body:
{
  "text": "Your video idea here"
}

Response:
{
  "scenario": "Enhanced script with delivery notations"
}
```

## Development

### Backend Development
```bash
cd backend

# Run tests
go test ./...

# Run with hot reload (install air first: go install github.com/cosmtrek/air@latest)
air

# Build for production
go build -o scenario-api main.go
```

### Frontend Development
```bash
cd web

# Development mode with hot reload
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

### Testing
```bash
# Backend tests
cd backend && go test ./... -v

# Test coverage
cd backend && go test ./... -cover
```

## Configuration

### Environment Variables
Create a `.env` file in the `backend/` directory:
```env
API_KEY=your_google_gemini_api_key
```

### Frontend Configuration
The frontend is configured to connect to the backend at `http://localhost:8080`. Update the API URL in `App.tsx` if needed.

## Features in Detail

### AI Enhancement
* Grammar and style improvement
* Natural, conversational tone
* Emotional cues and pacing suggestions
* Engagement optimization

### Error Handling
* Comprehensive request validation
* Graceful API failure handling
* User-friendly error messages
* Detailed server-side logging
* Rate limiting protection

## License

MIT — feel free to reuse and modify.
