import { useState, useEffect } from 'react'
import { TaskList } from './components/TaskList'
import { TaskForm } from './components/TaskForm'

const API_BASE = '/api/v1/tasks'

export default function App() {
  const [tasks, setTasks] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  const fetchTasks = async () => {
    setLoading(true)
    setError(null)
    try {
      const res = await fetch(API_BASE)
      if (!res.ok) throw new Error('Error al cargar tareas')
      const data = await res.json()
      setTasks(Array.isArray(data) ? data : [])
    } catch (e) {
      setError(e.message)
      setTasks([])
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchTasks()
  }, [])

  const handleCreated = () => {
    fetchTasks()
  }

  const handleStatusChange = () => {
    fetchTasks()
  }

  const handleDeleted = () => {
    fetchTasks()
  }

  return (
    <>
      <a href="#main-content" className="skip-link">
        Saltar al contenido principal
      </a>
      <header className="app-header" role="banner">
        <div className="container">
          <h1 className="app-title">Demo Perfil — Tasks</h1>
          <p className="app-desc">
            React · Diseño responsivo · Accesibilidad (WCAG)
          </p>
        </div>
      </header>
      <main id="main-content" className="app-main container" role="main" aria-label="Contenido principal">
        <section aria-labelledby="form-heading" className="section-form">
          <h2 id="form-heading" className="visually-hidden">Añadir nueva tarea</h2>
          <TaskForm onCreated={handleCreated} />
        </section>
        <section aria-labelledby="list-heading" className="section-list">
          <h2 id="list-heading">Lista de tareas</h2>
          {loading && (
            <p role="status" aria-live="polite">
              Cargando…
            </p>
          )}
          {error && (
            <div role="alert" className="error-banner">
              {error}
            </div>
          )}
          {!loading && !error && (
            <TaskList
              tasks={tasks}
              onStatusChange={handleStatusChange}
              onDeleted={handleDeleted}
            />
          )}
        </section>
      </main>
      <footer className="app-footer" role="contentinfo">
        <div className="container">
          <p>Backend: Go (Gin) · REST + GraphQL · PostgreSQL · Arquitectura Hexagonal</p>
        </div>
      </footer>
    </>
  )
}
