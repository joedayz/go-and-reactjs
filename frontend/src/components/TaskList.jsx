import { TaskCard } from './TaskCard'

export function TaskList({ tasks, onStatusChange, onDeleted }) {
  const list = Array.isArray(tasks) ? tasks : []
  if (list.length === 0) {
    return (
      <p className="empty-state" role="status">
        No hay tareas. Crea una desde el formulario.
      </p>
    )
  }
  return (
    <ul className="task-list" aria-label="Tareas">
      {list.map((task, index) => (
        <li key={task?.id ?? task?.ID ?? index}>
          <TaskCard
            task={task}
            onStatusChange={onStatusChange}
            onDeleted={onDeleted}
          />
        </li>
      ))}
    </ul>
  )
}
