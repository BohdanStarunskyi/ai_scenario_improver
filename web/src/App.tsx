import React, { useState } from 'react'
import './App.css'

function App() {
  const [idea, setIdea] = useState('')
  const [scenario, setScenario] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const [hasGenerated, setHasGenerated] = useState(false)

  const handleGenerate = async () => {
    if (!idea.trim()) return
    setIsLoading(true)
    
    const currentIdea = idea.trim()
    
    try {
      const response = await fetch('http://localhost:8080/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text: currentIdea }),
      })
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      const data = await response.json()
      const generatedScenario = data.scenario || 'No scenario generated.'
      
      setScenario(generatedScenario)
      setHasGenerated(true)
      
    } catch (error) {
      console.error('Error:', error)
      setScenario('An error occurred while generating the scenario. Please try again.')
      setHasGenerated(true)
    } finally {
      setIsLoading(false)
    }
  }

  const handleReset = () => {
    setIdea('')
    setScenario('')
    setHasGenerated(false)
  }

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !isLoading && !hasGenerated) {
      handleGenerate()
    }
  }

  return (
    <div className="app">
      <div className="background-blur"></div>
      
      <div className="container">
        <div className="header">
          <div className="logo">
            <div className="logo-icon">✨</div>
            <h1>ScenarioAI</h1>
          </div>
          <p className="subtitle">Transform your ideas into compelling scenarios</p>
        </div>

        <div className="main-content">
          {!hasGenerated ? (
            <div className="input-section">
              <div className="input-container">
                <div className="input-row">
                  <textarea
                    className="idea-input"
                    placeholder="Share your idea and watch it transform into an amazing scenario..."
                    value={idea}
                    onChange={(e) => setIdea(e.target.value)}
                    onKeyPress={handleKeyPress}
                    disabled={isLoading}
                    rows={4}
                  />
                  <div className="button-container">
                    <button
                      className="generate-btn"
                      onClick={handleGenerate}
                      disabled={isLoading || !idea.trim()}
                    >
                      {isLoading ? (
                        <div className="loading-spinner">
                          <div className="spinner"></div>
                          <span>Generating...</span>
                        </div>
                      ) : (
                        <>
                          <span>Generate Scenario</span>
                          <div className="btn-icon">🚀</div>
                        </>
                      )}
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ) : (
            <div className="result-section">
              <div className="idea-display">
                <div className="section-header">
                  <div className="section-icon">💡</div>
                  <h3>Your Idea</h3>
                </div>
                <div className="idea-content">{idea}</div>
              </div>

              <div className="scenario-display">
                <div className="section-header">
                  <div className="section-icon">🎬</div>
                  <h3>Generated Scenario</h3>
                </div>
                <div className="scenario-content">{scenario}</div>
              </div>

              <button className="reset-btn" onClick={handleReset}>
                <span>Create New Scenario</span>
                <div className="btn-icon">✨</div>
              </button>
            </div>
          )}
        </div>


      </div>
    </div>
  )
}

export default App