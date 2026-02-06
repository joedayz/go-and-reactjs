import { useState } from 'react'

const API_BASE = '/api/v1/tasks'

export function TaskForm({ onCreated }) {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState(null)

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError(null)
    if (!title.trim()) {
      setError('El título es obligatorio')
      return
    }
    setSubmitting(true)
    try {
      const res = await fetch(API_BASE, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          title: title.trim(),
          description: description.trim(),
        }),
      })
      if (!res.ok) {
        const data = await res.json().catch(() => ({}))
        throw new Error(data.error || 'Error al crear')
      }
      setTitle('')
      setDescription('')
      onCreated?.()
    } catch (e) {
      setError(e.message)
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <form
      onSubmit={handleSubmit}
      className="task-form"
      aria-label="Formulario nueva tarea"
      noValidate
    >
      {error && (
        <p role="alert" className="form-error">
          {error}
        </p>
      )}
      <div className="form-row">
        <label htmlFor="task-title">
          Título <span aria-hidden="true">*</span>
        </label>
        <input
          id="task-title"
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="Ej: Revisar documentación"
          disabled={submitting}
          autoComplete="off"
          aria-required="true"
          aria-invalid={!!error}
        />
      </div>
      <div className="form-row">
        <label htmlFor="task-desc">Descripción (opcional)</label>
        <textarea
          id="task-desc"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="Detalles opcionales"
          rows={2}
          disabled={submitting}
          aria-describedby="task-desc-hint"
        />
        <span id="task-desc-hint" className="hint">Opcional</span>
      </div>
      <button
        type="submit"
        className="btn btn-primary"
        disabled={submitting}
        aria-busy={submitting}
      >
        {submitting ? 'Creando…' : 'Crear tarea'}
      </button>
    </form>
  )
}
