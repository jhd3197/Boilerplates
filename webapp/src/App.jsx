import { useState, useEffect } from 'react'
import TemplateList from './components/TemplateList'
import TemplateForm from './components/TemplateForm'
import Header from './components/Header'
import { Loader2 } from 'lucide-react'

const REGISTRY_URL = 'https://raw.githubusercontent.com/jhd3197/boilerplates/main/templates-registry.json'

function App() {
  const [registry, setRegistry] = useState(null)
  const [selectedTemplate, setSelectedTemplate] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchRegistry()
  }, [])

  async function fetchRegistry() {
    try {
      setLoading(true)
      const response = await fetch(REGISTRY_URL)
      if (!response.ok) throw new Error('Failed to fetch registry')
      const data = await response.json()
      setRegistry(data)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  function handleSelectTemplate(template) {
    setSelectedTemplate(template)
  }

  function handleBack() {
    setSelectedTemplate(null)
  }

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-[#0d1117]">
        <Loader2 className="w-10 h-10 animate-spin text-[#1f6feb]" />
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-[#0d1117]">
        <div className="text-center">
          <p className="text-red-400 text-lg mb-4">Error: {error}</p>
          <button
            onClick={fetchRegistry}
            className="px-6 py-2 bg-[#238636] text-white rounded-md hover:bg-[#2ea043] transition-colors"
          >
            Retry
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen flex flex-col bg-[#0d1117]">
      <Header />
      {selectedTemplate ? (
        <TemplateForm
          template={selectedTemplate}
          registry={registry}
          onBack={handleBack}
        />
      ) : (
        <TemplateList
          templates={registry?.templates || []}
          onSelect={handleSelectTemplate}
        />
      )}
    </div>
  )
}

export default App
