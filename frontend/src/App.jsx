import { useState, useEffect } from 'react'

export default function App() {
  const [input, setInput] = useState('')
  const [options, setOptions] = useState([])

  useEffect(() => {
    fetch('/api/options')
      .then(r => r.json())
      .then(setOptions)
  }, [])

  const handleChange = async (e) => {
    const value = e.target.value
    setInput(value)
    const res = await fetch('/api/submit', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ value })
    })
    const data = await res.json()
    setOptions(data)
  }

  return (
    <div style={{ padding: '2rem', fontFamily: 'sans-serif' }}>
      <h1>PoC App</h1>
      <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem', maxWidth: '300px' }}>
        <input
          type="text"
          value={input}
          onChange={handleChange}
          placeholder="Type to filter..."
          style={{ padding: '0.5rem', fontSize: '1rem' }}
        />
        <select size={8} style={{ padding: '0.5rem', fontSize: '1rem' }}>
          {options.map(opt => (
            <option key={opt} value={opt}>{opt}</option>
          ))}
        </select>
      </div>
    </div>
  )
}
