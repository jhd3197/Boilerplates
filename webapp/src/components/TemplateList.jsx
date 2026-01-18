import { useState } from 'react'
import TemplateCard from './TemplateCard'
import { Search, Plus, Upload, Download } from 'lucide-react'

function TemplateList({ templates, onSelect }) {
  const [search, setSearch] = useState('')
  const [category, setCategory] = useState('all')

  const categories = ['all', ...new Set(templates.map(t => t.category))]

  const filtered = templates.filter(t => {
    const matchesSearch = t.name.toLowerCase().includes(search.toLowerCase()) ||
      t.description.toLowerCase().includes(search.toLowerCase()) ||
      t.tags?.some(tag => tag.toLowerCase().includes(search.toLowerCase()))
    const matchesCategory = category === 'all' || t.category === category
    return matchesSearch && matchesCategory
  })

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      {/* Control Bar */}
      <div className="px-8 py-6 flex justify-between items-center gap-4 flex-wrap">
        {/* Search & Filters */}
        <div className="flex items-center gap-4 flex-1">
          <div className="relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-[#8b949e]" />
            <input
              type="text"
              placeholder="Search templates..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              className="w-64 pl-10 pr-4 py-2 bg-[#0d1117] border border-[#30363d] rounded-md text-white text-sm focus:border-[#1f6feb] focus:outline-none transition-colors"
            />
          </div>

          <div className="flex gap-2">
            {categories.map(cat => (
              <button
                key={cat}
                onClick={() => setCategory(cat)}
                className={`px-4 py-1.5 rounded-full text-sm font-medium transition-all ${
                  category === cat
                    ? 'bg-[#c9d1d9] text-[#0d1117] font-semibold'
                    : 'bg-transparent border border-[#30363d] text-[#8b949e] hover:border-[#c9d1d9] hover:text-[#c9d1d9]'
                }`}
              >
                {cat.charAt(0).toUpperCase() + cat.slice(1)}
              </button>
            ))}
          </div>
        </div>

        {/* Action Buttons */}
        <div className="flex gap-3">
          <button className="flex items-center gap-2 px-4 py-2 bg-transparent border border-[#30363d] rounded-md text-[#c9d1d9] text-sm font-medium hover:bg-[#21262d] hover:border-[#8b949e] transition-all">
            <Upload className="w-4 h-4" />
            Import
          </button>
          <button className="flex items-center gap-2 px-4 py-2 bg-transparent border border-[#30363d] rounded-md text-[#c9d1d9] text-sm font-medium hover:bg-[#21262d] hover:border-[#8b949e] transition-all">
            <Download className="w-4 h-4" />
            Export
          </button>
          <button className="flex items-center gap-2 px-4 py-2 bg-[#238636] border border-white/10 rounded-md text-white text-sm font-medium hover:bg-[#2ea043] transition-all">
            <Plus className="w-4 h-4" />
            Add Boilerplate
          </button>
        </div>
      </div>

      {/* Template Grid */}
      <div className="flex-1 overflow-y-auto px-8 pb-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-[1400px] mx-auto">
          {filtered.map(template => (
            <TemplateCard
              key={template.id}
              template={template}
              onSelect={() => onSelect(template)}
            />
          ))}
        </div>

        {filtered.length === 0 && (
          <div className="text-center py-12">
            <p className="text-[#8b949e]">No templates found matching your criteria.</p>
          </div>
        )}
      </div>
    </div>
  )
}

export default TemplateList
