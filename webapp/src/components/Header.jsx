import { Package, Star } from 'lucide-react'

function Header() {
  return (
    <header className="bg-[#161b22] border-b border-[#30363d] px-8 py-4 flex justify-between items-center flex-shrink-0">
      <div className="flex items-center gap-4">
        <div className="w-8 h-8 bg-gradient-to-br from-[#1f6feb] to-[#238636] rounded-md flex items-center justify-center shadow-lg shadow-blue-500/20">
          <Package className="w-5 h-5 text-white" />
        </div>
        <span className="font-semibold text-white text-lg">Boilerplate Manager</span>
      </div>

      <a
        href="https://github.com/jhd3197/boilerplates"
        target="_blank"
        rel="noopener noreferrer"
        className="flex items-center gap-2 px-3 py-1.5 bg-[#21262d] border border-[#30363d] rounded-md text-[#c9d1d9] text-sm font-semibold hover:bg-[#30363d] hover:border-[#8b949e] hover:text-white transition-all"
      >
        <Star className="w-4 h-4 fill-[#e3b341] text-[#e3b341]" />
        Star on GitHub
      </a>
    </header>
  )
}

export default Header
