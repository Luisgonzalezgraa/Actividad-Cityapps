import { ref } from 'vue'

export type Task = {
  _id: string
  title: string
  description: string
  completed: boolean
  active: boolean
  isDeleted: boolean
  tags: string[]
  createdAt: string
  updatedAt: string
}

export const useTasks = () => {
  const config = useRuntimeConfig()
  const tasks = ref<Task[]>([])
  const loading = ref(false)
  const error = ref<string>('')

  const fetchTasks = async () => {
    try {
      loading.value = true
      error.value = ''
      const response = await $fetch<Task[]>(`${config.public.apiBase}/tasks`)
      tasks.value = response || []
    } catch (err: any) {
      error.value = err?.message || 'No se pudieron cargar las tareas'
    } finally {
      loading.value = false
    }
  }

  const createTask = async (payload: { title: string; description: string; tags: string[] }) => {
    try {
      await $fetch(`${config.public.apiBase}/tasks`, {
        method: 'POST',
        body: payload
      })
      await fetchTasks()
    } catch (err: any) {
      error.value = err?.message || 'Error al crear la tarea'
    }
  }

  const completeTask = async (id: string) => {
    try {
      await $fetch(`${config.public.apiBase}/tasks/${id}/complete`, {
        method: 'PUT'
      })
      await fetchTasks()
    } catch (err: any) {
      error.value = err?.message || 'Error al completar la tarea'
    }
  }

  const deleteTask = async (id: string) => {
    try {
      await $fetch(`${config.public.apiBase}/tasks/${id}`, {
        method: 'DELETE'
      })
      await fetchTasks()
    } catch (err: any) {
      error.value = err?.message || 'Error al eliminar la tarea'
    }
  }

  const deletePermanent = async (id: string) => {
    try {
      await $fetch(`${config.public.apiBase}/tasks/${id}/permanent`, {
        method: 'DELETE'
      })
      await fetchTasks()
    } catch (err: any) {
      error.value = err?.message || 'Error al eliminar permanentemente la tarea'
    }
  }

  const recoverTask = async (id: string) => {
    try {
      await $fetch(`${config.public.apiBase}/tasks/${id}/recover`, {
        method: 'PUT'
      })
      await fetchTasks()
    } catch (err: any) {
      error.value = err?.message || 'Error al recuperar la tarea'
    }
  }

  const toggleActive = async (id: string) => {
    try {
      const task = tasks.value.find(t => t._id === id)
      if (!task) return
      
      await $fetch(`${config.public.apiBase}/tasks/${id}/toggle-active`, {
        method: 'PUT'
      })
      await fetchTasks()
    } catch (err: any) {
      error.value = err?.message || 'Error al cambiar estado de actividad'
    }
  }

  return {
    tasks,
    loading,
    error,
    fetchTasks,
    createTask,
    completeTask,
    deleteTask,
    deletePermanent,
    recoverTask,
    toggleActive
  }
}