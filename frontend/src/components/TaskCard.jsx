import { useState } from 'react'

const API_BASE = '/api/v1/tasks'
const STATUS_LABELS = {
  PENDING: 'Pendiente',
  IN_PROGRESS: 'En progreso',
  DONE: 'Hecho',
}

export function TaskCard({ task, onStatusChange, onDeleted }) {
  const [updating, setUpdating] = useState(false)
  const [deleting, setDeleting] = useState(false)
  const id = task?.id ?? task?.ID ?? ''
  const title = task?.title ?? task?.Title ?? ''
  const description = task?.description ?? task?.Description ?? ''
  const status = task?.status ?? task?.Status ?? 'PENDING'

  const updateStatus = async (newStatus) => {
    if (!id) return
    setUpdating(true)
    try {
      const res = await fetch(`${API_BASE}/${id}/status`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ status: newStatus }),
      })
      if (res.ok) onStatusChange?.()
    } finally {
      setUpdating(false)
    }
  }

  const deleteTask = async () => {
    if (!id) return
    if (!window.confirm('¿Eliminar esta tarea?')) return
    setDeleting(true)
    try {
      const res = await fetch(`${API_BASE}/${id}`, { method: 'DELETE' })
      if (res.ok) onDeleted?.()
    } finally {
      setDeleting(false)
    }
  }

  const statusId = `status-${id}`
  const cardId = `task-${id}`

  return (
    <article
      id={cardId}
      className="task-card"
      aria-labelledby={`${cardId}-title`}
      aria-describedby={description ? `${cardId}-desc` : undefined}
    >
      <div className="task-card-header">
        <h3 id={`${cardId}-title`} className="task-title">
          {title}
        </h3>
        <span
          id={statusId}
          className={`task-status task-status--${String(status).toLowerCase().replace('_', '-')}`}
          aria-label={`Estado: ${STATUS_LABELS[status] || status}`}
        >
          {STATUS_LABELS[status] || status}
        </span>
      </div>
      {description && (
        <p id={`${cardId}-desc`} className="task-description">
          {description}
        </p>
      )}
      <div className="task-card-actions">
        <label htmlFor={`select-${id}`} className="visually-hidden">
          Cambiar estado de la tarea
        </label>
        <select
          id={`select-${id}`}
          value={status}
          onChange={(e) => updateStatus(e.target.value)}
          disabled={updating}
          aria-labelledby={statusId}
        >
          {Object.entries(STATUS_LABELS).map(([value, label]) => (
            <option key={value} value={value}>
              {label}
            </option>
          ))}
        </select>
        <button
          type="button"
          className="btn btn-ghost btn-danger"
          onClick={deleteTask}
          disabled={deleting}
          aria-label={`Eliminar tarea "${title}"`}
        >
          {deleting ? '…' : 'Eliminar'}
        </button>
      </div>
    </article>
  )
}
