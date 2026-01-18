import { useState, useEffect } from 'react'
import { ArrowLeft, Download, Loader2, FolderTree } from 'lucide-react'
import { generateZip } from '../utils/zipGenerator'

function TemplateForm({ template, registry, onBack }) {
  const [templateConfig, setTemplateConfig] = useState(null)
  const [values, setValues] = useState({})
  const [loading, setLoading] = useState(true)
  const [downloading, setDownloading] = useState(false)
  const [error, setError] = useState(null)

  const repoUrl = template.repo || registry.default_repo
  const branch = template.branch || 'main'

  useEffect(() => {
    fetchTemplateConfig()
  }, [template])

  async function fetchTemplateConfig() {
    try {
      setLoading(true)
      const configUrl = `https://raw.githubusercontent.com/${getRepoPath(repoUrl)}/${branch}/${template.path}/template.json`
      const response = await fetch(configUrl)
      if (!response.ok) throw new Error('Failed to fetch template config')
      const config = await response.json()
      setTemplateConfig(config)

      const initialValues = {}
      if (config.prompts) {
        Object.entries(config.prompts).forEach(([key, prompt]) => {
          initialValues[key] = prompt.default || ''
        })
      }
      setValues(initialValues)
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  function getRepoPath(url) {
    const match = url.match(/github\.com\/([^\/]+\/[^\/]+)/)
    return match ? match[1] : url
  }

  function handleChange(key, value) {
    setValues(prev => ({ ...prev, [key]: value }))
  }

  function formatValue(value, format) {
    if (!format || format === 'text') return value
    if (format === 'snake_case') {
      return value.toLowerCase().replace(/\s+/g, '_').replace(/[^a-z0-9_]/g, '')
    }
    return value
  }

  async function handleDownload() {
    try {
      setDownloading(true)
      setError(null)

      const formattedValues = {}
      if (templateConfig.prompts) {
        Object.entries(templateConfig.prompts).forEach(([key, prompt]) => {
          formattedValues[key] = formatValue(values[key], prompt.format)
        })
      }

      await generateZip({
        repoUrl,
        branch,
        templatePath: template.path,
        templateConfig,
        values: formattedValues,
        projectName: values.project_name || values.package_name || template.id,
      })
    } catch (err) {
      setError(err.message)
    } finally {
      setDownloading(false)
    }
  }

  if (loading) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <Loader2 className="w-8 h-8 animate-spin text-[#1f6feb]" />
      </div>
    )
  }

  return (
    <div className="flex-1 overflow-y-auto px-8 py-6">
      <div className="max-w-2xl mx-auto">
        {/* Back Button */}
        <button
          onClick={onBack}
          className="flex items-center gap-2 text-[#8b949e] hover:text-white mb-6 transition-colors"
        >
          <ArrowLeft className="w-4 h-4" />
          Back to templates
        </button>

        {/* Template Header */}
        <div className="bg-[#161b22] border border-[#30363d] rounded-md p-6 mb-6">
          <h2 className="text-2xl font-bold text-white mb-2">{template.name}</h2>
          <p className="text-[#8b949e] mb-4">{template.description}</p>
          <div className="flex items-center gap-2 text-sm text-[#8b949e]">
            <FolderTree className="w-4 h-4" />
            <span>{template.path}</span>
          </div>
        </div>

        {/* Form */}
        <div className="bg-[#161b22] border border-[#30363d] rounded-md p-6">
          <h3 className="text-lg font-semibold text-white mb-6">Configure Your Project</h3>

          {templateConfig?.prompts && Object.entries(templateConfig.prompts).map(([key, prompt]) => (
            <div key={key} className="mb-5">
              <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
                {prompt.label}
              </label>
              <input
                type="text"
                value={values[key] || ''}
                onChange={(e) => handleChange(key, e.target.value)}
                placeholder={prompt.default}
                className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors"
              />
              {prompt.format && prompt.format !== 'text' && (
                <p className="text-xs text-[#8b949e] mt-1.5">Format: {prompt.format}</p>
              )}
            </div>
          ))}

          {error && (
            <div className="mb-5 p-4 bg-red-500/10 border border-red-500/20 text-red-400 rounded-md text-sm">
              {error}
            </div>
          )}

          <button
            onClick={handleDownload}
            disabled={downloading}
            className="w-full flex items-center justify-center gap-2 px-6 py-3 bg-[#238636] border border-white/10 text-white rounded-md font-semibold hover:bg-[#2ea043] disabled:opacity-50 disabled:cursor-not-allowed transition-all"
          >
            {downloading ? (
              <>
                <Loader2 className="w-5 h-5 animate-spin" />
                Generating...
              </>
            ) : (
              <>
                <Download className="w-5 h-5" />
                Download Project
              </>
            )}
          </button>
        </div>
      </div>
    </div>
  )
}

export default TemplateForm
