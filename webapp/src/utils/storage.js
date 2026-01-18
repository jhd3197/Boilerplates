const STORAGE_KEY = 'boilerplate-manager-custom-templates'

export function getCustomTemplates() {
  try {
    const stored = localStorage.getItem(STORAGE_KEY)
    return stored ? JSON.parse(stored) : []
  } catch (error) {
    console.error('Error reading custom templates from localStorage:', error)
    return []
  }
}

export function saveCustomTemplates(templates) {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(templates))
  } catch (error) {
    console.error('Error saving custom templates to localStorage:', error)
  }
}

export function addCustomTemplate(template) {
  const templates = getCustomTemplates()
  const existingIndex = templates.findIndex(t => t.id === template.id)

  if (existingIndex >= 0) {
    templates[existingIndex] = template
  } else {
    templates.push(template)
  }

  saveCustomTemplates(templates)
  return templates
}

export function updateCustomTemplate(template) {
  const templates = getCustomTemplates()
  const index = templates.findIndex(t => t.id === template.id)

  if (index >= 0) {
    templates[index] = template
    saveCustomTemplates(templates)
  }

  return templates
}

export function deleteCustomTemplate(templateId) {
  const templates = getCustomTemplates()
  const filtered = templates.filter(t => t.id !== templateId)
  saveCustomTemplates(filtered)
  return filtered
}

export function exportTemplates() {
  const templates = getCustomTemplates()
  const blob = new Blob([JSON.stringify(templates, null, 2)], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = 'custom-templates.json'
  a.click()
  URL.revokeObjectURL(url)
}

export function importTemplates(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = (e) => {
      try {
        const imported = JSON.parse(e.target.result)
        if (!Array.isArray(imported)) {
          throw new Error('Invalid format: expected an array of templates')
        }

        const templates = getCustomTemplates()
        let added = 0
        let updated = 0

        imported.forEach(template => {
          if (!template.id || !template.name) {
            return // Skip invalid templates
          }

          template.isCustom = true
          const existingIndex = templates.findIndex(t => t.id === template.id)

          if (existingIndex >= 0) {
            templates[existingIndex] = template
            updated++
          } else {
            templates.push(template)
            added++
          }
        })

        saveCustomTemplates(templates)
        resolve({ added, updated, total: templates.length })
      } catch (error) {
        reject(new Error('Failed to parse import file: ' + error.message))
      }
    }
    reader.onerror = () => reject(new Error('Failed to read file'))
    reader.readAsText(file)
  })
}
