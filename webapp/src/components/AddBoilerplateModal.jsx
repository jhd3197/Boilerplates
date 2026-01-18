import { useState } from 'react'
import { Plus, GitBranch, FolderTree, Lock, Tag } from 'lucide-react'
import Modal from './Modal'

function AddBoilerplateModal({ isOpen, onClose, onAdd }) {
  const [formData, setFormData] = useState({
    id: '',
    name: '',
    description: '',
    repo: '',
    path: '',
    branch: 'main',
    commit: '',
    category: 'other',
    tags: '',
    isPrivate: false,
  })
  const [errors, setErrors] = useState({})

  function handleChange(field, value) {
    setFormData(prev => ({ ...prev, [field]: value }))
    if (errors[field]) {
      setErrors(prev => ({ ...prev, [field]: null }))
    }
  }

  function validateForm() {
    const newErrors = {}

    if (!formData.id.trim()) {
      newErrors.id = 'Template ID is required'
    } else if (!/^[a-z0-9_-]+$/.test(formData.id)) {
      newErrors.id = 'ID must be lowercase letters, numbers, hyphens, or underscores'
    }

    if (!formData.name.trim()) {
      newErrors.name = 'Name is required'
    }

    if (!formData.repo.trim()) {
      newErrors.repo = 'Repository URL is required'
    } else if (!formData.repo.includes('github.com')) {
      newErrors.repo = 'Must be a valid GitHub repository URL'
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  function handleSubmit(e) {
    e.preventDefault()

    if (!validateForm()) return

    const template = {
      id: formData.id.trim(),
      name: formData.name.trim(),
      description: formData.description.trim(),
      repo: formData.repo.trim(),
      path: formData.path.trim() || '.',
      branch: formData.branch.trim() || 'main',
      commit: formData.commit.trim() || null,
      category: formData.category,
      tags: formData.tags.split(',').map(t => t.trim()).filter(Boolean),
      isPrivate: formData.isPrivate,
      isCustom: true,
    }

    onAdd(template)
    handleReset()
    onClose()
  }

  function handleReset() {
    setFormData({
      id: '',
      name: '',
      description: '',
      repo: '',
      path: '',
      branch: 'main',
      commit: '',
      category: 'other',
      tags: '',
      isPrivate: false,
    })
    setErrors({})
  }

  return (
    <Modal isOpen={isOpen} onClose={onClose} title="Add Boilerplate">
      <form onSubmit={handleSubmit} className="space-y-5">
        {/* Template ID */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Template ID <span className="text-red-400">*</span>
          </label>
          <input
            type="text"
            value={formData.id}
            onChange={(e) => handleChange('id', e.target.value.toLowerCase())}
            placeholder="my-template"
            className={`w-full px-4 py-2.5 bg-[#0d1117] border rounded-md text-white focus:outline-none transition-colors ${
              errors.id ? 'border-red-500 focus:border-red-500' : 'border-[#30363d] focus:border-[#1f6feb]'
            }`}
          />
          {errors.id && <p className="text-red-400 text-xs mt-1">{errors.id}</p>}
        </div>

        {/* Name */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Display Name <span className="text-red-400">*</span>
          </label>
          <input
            type="text"
            value={formData.name}
            onChange={(e) => handleChange('name', e.target.value)}
            placeholder="My Awesome Template"
            className={`w-full px-4 py-2.5 bg-[#0d1117] border rounded-md text-white focus:outline-none transition-colors ${
              errors.name ? 'border-red-500 focus:border-red-500' : 'border-[#30363d] focus:border-[#1f6feb]'
            }`}
          />
          {errors.name && <p className="text-red-400 text-xs mt-1">{errors.name}</p>}
        </div>

        {/* Description */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Description
          </label>
          <textarea
            value={formData.description}
            onChange={(e) => handleChange('description', e.target.value)}
            placeholder="A brief description of your template..."
            rows={2}
            className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors resize-none"
          />
        </div>

        {/* Repository URL */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            <span className="flex items-center gap-2">
              GitHub Repository URL <span className="text-red-400">*</span>
            </span>
          </label>
          <input
            type="text"
            value={formData.repo}
            onChange={(e) => handleChange('repo', e.target.value)}
            placeholder="https://github.com/username/repo"
            className={`w-full px-4 py-2.5 bg-[#0d1117] border rounded-md text-white focus:outline-none transition-colors ${
              errors.repo ? 'border-red-500 focus:border-red-500' : 'border-[#30363d] focus:border-[#1f6feb]'
            }`}
          />
          {errors.repo && <p className="text-red-400 text-xs mt-1">{errors.repo}</p>}
        </div>

        {/* Path and Branch - Side by side */}
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
              <span className="flex items-center gap-2">
                <FolderTree className="w-4 h-4" />
                Path in Repo
              </span>
            </label>
            <input
              type="text"
              value={formData.path}
              onChange={(e) => handleChange('path', e.target.value)}
              placeholder="templates/my-template"
              className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
              <span className="flex items-center gap-2">
                <GitBranch className="w-4 h-4" />
                Branch
              </span>
            </label>
            <input
              type="text"
              value={formData.branch}
              onChange={(e) => handleChange('branch', e.target.value)}
              placeholder="main"
              className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors"
            />
          </div>
        </div>

        {/* Commit Pin */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Pin to Commit (optional)
          </label>
          <input
            type="text"
            value={formData.commit}
            onChange={(e) => handleChange('commit', e.target.value)}
            placeholder="abc123... (leave empty for latest)"
            className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors font-mono text-sm"
          />
        </div>

        {/* Category and Tags - Side by side */}
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
              Category
            </label>
            <select
              value={formData.category}
              onChange={(e) => handleChange('category', e.target.value)}
              className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors"
            >
              <option value="react">React</option>
              <option value="python">Python</option>
              <option value="php">PHP</option>
              <option value="node">Node.js</option>
              <option value="go">Go</option>
              <option value="other">Other</option>
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
              <span className="flex items-center gap-2">
                <Tag className="w-4 h-4" />
                Tags
              </span>
            </label>
            <input
              type="text"
              value={formData.tags}
              onChange={(e) => handleChange('tags', e.target.value)}
              placeholder="api, docker, auth"
              className="w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors"
            />
          </div>
        </div>

        {/* Private Repo Checkbox */}
        <div className="flex items-center gap-3">
          <input
            type="checkbox"
            id="isPrivate"
            checked={formData.isPrivate}
            onChange={(e) => handleChange('isPrivate', e.target.checked)}
            className="w-4 h-4 rounded border-[#30363d] bg-[#0d1117] text-[#238636] focus:ring-[#238636] focus:ring-offset-0"
          />
          <label htmlFor="isPrivate" className="flex items-center gap-2 text-sm text-[#c9d1d9]">
            <Lock className="w-4 h-4" />
            Private repository (requires GitHub token)
          </label>
        </div>

        {/* Actions */}
        <div className="flex gap-3 pt-4 border-t border-[#30363d]">
          <button
            type="button"
            onClick={() => { handleReset(); onClose(); }}
            className="flex-1 py-2.5 px-4 bg-transparent border border-[#30363d] rounded-md text-[#c9d1d9] font-medium hover:bg-[#21262d] transition-colors"
          >
            Cancel
          </button>
          <button
            type="submit"
            className="flex-1 flex items-center justify-center gap-2 py-2.5 px-4 bg-[#238636] border border-white/10 rounded-md text-white font-medium hover:bg-[#2ea043] transition-colors"
          >
            <Plus className="w-4 h-4" />
            Add Template
          </button>
        </div>
      </form>
    </Modal>
  )
}

export default AddBoilerplateModal
