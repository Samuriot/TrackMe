import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import './index.css'
import App from './App.tsx'

const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)")

function applySystemTheme() {
  document.documentElement.classList.toggle("dark", mediaQuery.matches)
}

applySystemTheme()
mediaQuery.addEventListener("change", applySystemTheme)

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </StrictMode>,
)
