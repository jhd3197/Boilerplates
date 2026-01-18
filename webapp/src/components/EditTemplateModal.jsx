import { useState, useEffect } from 'react'
import { Save, Trash2, GitBranch, FolderTree, Lock, Tag, ExternalLink } from 'lucide-react'
import Modal from './Modal'

function EditTemplateModal({ isOpen, onClose, template, onSave, onDelete }) {
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
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false)

  useEffect(() => {
    if (template) {
      setFormData({
        id: template.id || '',
        name: template.name || '',
        description: template.description || '',
        repo: template.repo || '',
        path: template.path || '',
        branch: template.branch || 'main',
        commit: template.commit || '',
        category: template.category || 'other',
        tags: Array.isArray(template.tags) ? template.tags.join(', ') : '',
        isPrivate: template.isPrivate || false,
      })
    }
  }, [template])

  function handleChange(field, value) {
    setFormData(prev => ({ ...prev, [field]: value }))
  }

  function handleSubmit(e) {
    e.preventDefault()

    const updatedTemplate = {
      ...template,
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
    }

    onSave(updatedTemplate)
    onClose()
  }

  function handleDelete() {
    onDelete(template.id)
    onClose()
  }

  const isCustomTemplate = template?.isCustom

  return (
    <Modal isOpen={isOpen} onClose={onClose} title={isCustomTemplate ? "Edit Template" : "Template Details"}>
      <form onSubmit={handleSubmit} className="space-y-5">
        {/* Info banner for registry templates */}
        {!isCustomTemplate && (
          <div className="p-3 bg-[#1f6feb]/10 border border-[#1f6feb]/20 rounded-md text-sm text-[#58a6ff]">
            This is a registry template. You can view its details but cannot edit it.
            To customize, add it as a custom template.
          </div>
        )}

        {/* Template ID */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Template ID
          </label>
          <input
            type="text"
            value={formData.id}
            onChange={(e) => handleChange('id', e.target.value.toLowerCase())}
            disabled={!isCustomTemplate}
            className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
              !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
            }`}
          />
        </div>

        {/* Name */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Display Name
          </label>
          <input
            type="text"
            value={formData.name}
            onChange={(e) => handleChange('name', e.target.value)}
            disabled={!isCustomTemplate}
            className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
              !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
            }`}
          />
        </div>

        {/* Description */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Description
          </label>
          <textarea
            value={formData.description}
            onChange={(e) => handleChange('description', e.target.value)}
            disabled={!isCustomTemplate}
            rows={2}
            className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors resize-none ${
              !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
            }`}
          />
        </div>

        {/* Repository URL */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            GitHub Repository
          </label>
          <div className="flex gap-2">
            <input
              type="text"
              value={formData.repo}
              onChange={(e) => handleChange('repo', e.target.value)}
              disabled={!isCustomTemplate}
              className={`flex-1 px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
                !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
              }`}
            />
            {formData.repo && (
              <a
                href={formData.repo}
                target="_blank"
                rel="noopener noreferrer"
                className="px-3 py-2.5 bg-[#21262d] border border-[#30363d] rounded-md text-[#c9d1d9] hover:bg-[#30363d] hover:text-white transition-colors"
              >
                <ExternalLink className="w-4 h-4" />
              </a>
            )}
          </div>
        </div>

        {/* Path and Branch */}
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
              disabled={!isCustomTemplate}
              className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
                !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
              }`}
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
              disabled={!isCustomTemplate}
              className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
                !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
              }`}
            />
          </div>
        </div>

        {/* Commit Pin */}
        <div>
          <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
            Pinned Commit
          </label>
          <input
            type="text"
            value={formData.commit || ''}
            onChange={(e) => handleChange('commit', e.target.value)}
            disabled={!isCustomTemplate}
            placeholder={isCustomTemplate ? "Leave empty for latest" : "Not pinned (using latest)"}
            className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors font-mono text-sm ${
              !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
            }`}
          />
        </div>

        {/* Category and Tags */}
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-[#c9d1d9] mb-2">
              Category
            </label>
            <select
              value={formData.category}
              onChange={(e) => handleChange('category', e.target.value)}
              disabled={!isCustomTemplate}
              className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
                !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
              }`}
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
              disabled={!isCustomTemplate}
              className={`w-full px-4 py-2.5 bg-[#0d1117] border border-[#30363d] rounded-md text-white focus:border-[#1f6feb] focus:outline-none transition-colors ${
                !isCustomTemplate ? 'opacity-60 cursor-not-allowed' : ''
              }`}
            />
          </div>
        </div>

        {/* Private Repo Checkbox */}
        {isCustomTemplate && (
          <div className="flex items-center gap-3">
            <input
              type="checkbox"
              id="editIsPrivate"
              checked={formData.isPrivate}
              onChange={(e) => handleChange('isPrivate', e.target.checked)}
              className="w-4 h-4 rounded border-[#30363d] bg-[#0d1117] text-[#238636] focus:ring-[#238636] focus:ring-offset-0"
            />
            <label htmlFor="editIsPrivate" className="flex items-center gap-2 text-sm text-[#c9d1d9]">
              <Lock className="w-4 h-4" />
              Private repository
            </label>
          </div>
        )}

        {/* Actions */}
        <div className="flex gap-3 pt-4 border-t border-[#30363d]">
          {isCustomTemplate ? (
            <>
              {showDeleteConfirm ? (
                <>
                  <button
                    type="button"
                    onClick={() => setShowDeleteConfirm(false)}
                    className="flex-1 py-2.5 px-4 bg-transparent border border-[#30363d] rounded-md text-[#c9d1d9] font-medium hover:bg-[#21262d] transition-colors"
                  >
                    Cancel
                  </button>
                  <button
                    type="button"
                    onClick={handleDelete}
                    className="flex-1 flex items-center justify-center gap-2 py-2.5 px-4 bg-red-600 border border-red-500 rounded-md text-white font-medium hover:bg-red-700 transition-colors"
                  >
                    <Trash2 className="w-4 h-4" />
                    Confirm Delete
                  </button>
                </>
              ) : (
                <>
                  <button
                    type="button"
                    onClick={() => setShowDeleteConfirm(true)}
                    className="py-2.5 px-4 bg-transparent border border-red-500/50 rounded-md text-red-400 font-medium hover:bg-red-500/10 transition-colors"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                  <button
                    type="button"
                    onClick={onClose}
                    className="flex-1 py-2.5 px-4 bg-transparent border border-[#30363d] rounded-md text-[#c9d1d9] font-medium hover:bg-[#21262d] transition-colors"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    className="flex-1 flex items-center justify-center gap-2 py-2.5 px-4 bg-[#238636] border border-white/10 rounded-md text-white font-medium hover:bg-[#2ea043] transition-colors"
                  >
                    <Save className="w-4 h-4" />
                    Save Changes
                  </button>
                </>
              )}
            </>
          ) : (
            <button
              type="button"
              onClick={onClose}
              className="flex-1 py-2.5 px-4 bg-[#21262d] border border-[#30363d] rounded-md text-[#c9d1d9] font-medium hover:bg-[#30363d] transition-colors"
            >
              Close
            </button>
          )}
        </div>
      </form>
    </Modal>
  )
}

export default EditTemplateModal
