import { Pencil } from 'lucide-react'

const categoryIcons = {
  react: '⚛️',
  python: '🐍',
  php: '🐘',
  other: '📦',
}

function TemplateCard({ template, onSelect, onEdit }) {
  const icon = categoryIcons[template.category] || categoryIcons.other

  return (
    <div className="bg-[#161b22] border border-[#30363d] rounded-md p-5 flex flex-col transition-all duration-200 hover:border-[#8b949e] hover:-translate-y-1 hover:shadow-xl hover:shadow-black/30 relative z-0 hover:z-10">
      <div className="flex justify-between items-center mb-4">
        <h3 className="text-white font-semibold text-lg flex items-center gap-2">
          <span>{icon}</span>
          {template.name}
        </h3>
      </div>

      <p className="text-[#8b949e] text-sm mb-6 leading-relaxed flex-grow">
        {template.description}
      </p>

      {template.tags && template.tags.length > 0 && (
        <div className="flex flex-wrap gap-2 mb-6">
          {template.tags.slice(0, 4).map(tag => (
            <span
              key={tag}
              className="text-xs text-[#1f6feb] bg-[#1f6feb]/10 border border-[#1f6feb]/20 px-2 py-0.5 rounded-full"
            >
              {tag}
            </span>
          ))}
        </div>
      )}

      <div className="border-t border-[#30363d] pt-4 flex gap-2">
        <button
          onClick={(e) => {
            e.stopPropagation()
            onEdit()
          }}
          className="flex-1 py-2 px-4 rounded bg-transparent border border-[#30363d] text-[#c9d1d9] text-sm hover:bg-[#21262d] hover:text-white transition-all"
        >
          <Pencil className="w-4 h-4 inline mr-2" />
          Edit
        </button>
        <button
          onClick={onSelect}
          className="flex-1 py-2 px-4 rounded bg-[#1f6feb] border border-[#1f6feb] text-white text-sm font-semibold hover:bg-[#3b82f6] transition-all"
        >
          Create Project
        </button>
      </div>
    </div>
  )
}

export default TemplateCard
