import { useState, useEffect } from 'react'
import TemplateList from './components/TemplateList'
import TemplateForm from './components/TemplateForm'
import Header from './components/Header'
import AddBoilerplateModal from './components/AddBoilerplateModal'
import EditTemplateModal from './components/EditTemplateModal'
import { Loader2 } from 'lucide-react'
import {
  getCustomTemplates,
  addCustomTemplate,
  updateCustomTemplate,
  deleteCustomTemplate,
  exportTemplates,
  importTemplates,
} from './utils/storage'

const REGISTRY_URL = 'https://raw.githubusercontent.com/jhd3197/boilerplates/main/templates-registry.json'

function App() {
  const [registry, setRegistry] = useState(null)
  const [customTemplates, setCustomTemplates] = useState([])
  const [selectedTemplate, setSelectedTemplate] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  // Modal states
  const [showAddModal, setShowAddModal] = useState(false)
  const [editingTemplate, setEditingTemplate] = useState(null)

  useEffect(() => {
    fetchRegistry()
    setCustomTemplates(getCustomTemplates())
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

  // Combine registry templates with custom templates
  const allTemplates = [
    ...(registry?.templates || []),
    ...customTemplates,
  ]

  function handleSelectTemplate(template) {
    setSelectedTemplate(template)
  }

  function handleBack() {
    setSelectedTemplate(null)
  }

  function handleAddTemplate(template) {
    const updated = addCustomTemplate(template)
    setCustomTemplates(updated)
  }

  function handleEditTemplate(template) {
    setEditingTemplate(template)
  }

  function handleSaveTemplate(template) {
    const updated = updateCustomTemplate(template)
    setCustomTemplates(updated)
  }

  function handleDeleteTemplate(templateId) {
    const updated = deleteCustomTemplate(templateId)
    setCustomTemplates(updated)
  }

  function handleExport() {
    exportTemplates()
  }

  async function handleImport() {
    const input = document.createElement('input')
    input.type = 'file'
    input.accept = '.json'
    input.onchange = async (e) => {
      const file = e.target.files?.[0]
      if (!file) return

      try {
        const result = await importTemplates(file)
        setCustomTemplates(getCustomTemplates())
        alert(`Import successful!\nAdded: ${result.added}\nUpdated: ${result.updated}`)
      } catch (err) {
        alert('Import failed: ' + err.message)
      }
    }
    input.click()
  }

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-[#0d1117]">
        <Loader2 className="w-10 h-10 animate-spin text-[#1f6feb]" />
      </div>
    )
  }

  if (error && !registry) {
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
          templates={allTemplates}
          onSelect={handleSelectTemplate}
          onEdit={handleEditTemplate}
          onAdd={() => setShowAddModal(true)}
          onExport={handleExport}
          onImport={handleImport}
        />
      )}

      {/* Add Boilerplate Modal */}
      <AddBoilerplateModal
        isOpen={showAddModal}
        onClose={() => setShowAddModal(false)}
        onAdd={handleAddTemplate}
      />

      {/* Edit Template Modal */}
      <EditTemplateModal
        isOpen={!!editingTemplate}
        onClose={() => setEditingTemplate(null)}
        template={editingTemplate}
        onSave={handleSaveTemplate}
        onDelete={handleDeleteTemplate}
      />
    </div>
  )
}

export default App
