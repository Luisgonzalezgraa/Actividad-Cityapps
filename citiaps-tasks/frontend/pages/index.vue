<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useTasks } from '~/composables/useTasks'
import TaskForm from '~/components/TaskForm.vue'
import TaskList from '~/components/TaskList.vue'

const { tasks, loading, error, fetchTasks, createTask, completeTask, deleteTask, deletePermanent, recoverTask, toggleActive } = useTasks()

const filterType = ref<'all' | 'active' | 'completed' | 'inactive' | 'deleted'>('all')
const currentPage = ref(1)
const pageSize = 5
const showDeleteModal = ref(false)
const taskToDelete = ref<string | null>(null)
const showDeletePermanentModal = ref(false)
const taskToDeletePermanent = ref<string | null>(null)

const filteredTasks = computed(() => {
  const activeTasks = tasks.value.filter(t => !t.isDeleted && t.active)
  const inactiveTasks = tasks.value.filter(t => !t.isDeleted && !t.active)
  const completedTasks = tasks.value.filter(t => !t.isDeleted && t.completed)
  const deletedTasks = tasks.value.filter(t => t.isDeleted)

  if (filterType.value === 'deleted') {
    return deletedTasks
  }

  if (filterType.value === 'active') {
    return activeTasks.filter(t => !t.completed)
  }
  if (filterType.value === 'completed') {
    return completedTasks
  }
  if (filterType.value === 'inactive') {
    return inactiveTasks
  }
  return tasks.value.filter(t => !t.isDeleted)
})

// Paginación: aplicar después de filtrar
const paginatedTasks = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  const end = start + pageSize
  return filteredTasks.value.slice(start, end)
})

const totalPages = computed(() => {
  return Math.ceil(filteredTasks.value.length / pageSize)
})

const pendingCount = computed(() => tasks.value.filter(t => !t.isDeleted && !t.completed && t.active).length)
const completedCount = computed(() => tasks.value.filter(t => !t.isDeleted && t.completed).length)
const inactiveCount = computed(() => tasks.value.filter(t => !t.isDeleted && !t.active).length)
const deletedCount = computed(() => tasks.value.filter(t => t.isDeleted).length)

const confirmDelete = (id: string) => {
  taskToDelete.value = id
  showDeleteModal.value = true
}

const handleDelete = async () => {
  if (taskToDelete.value) {
    await deleteTask(taskToDelete.value)
    showDeleteModal.value = false
    taskToDelete.value = null
  }
}

const handleRecover = async (id: string) => {
  await recoverTask(id)
}

const handleToggleActive = async (id: string) => {
  await toggleActive(id)
}

const handleRemovePermanent = async (id: string) => {
  taskToDeletePermanent.value = id
  showDeletePermanentModal.value = true
}

const handleDeletePermanent = async () => {
  if (taskToDeletePermanent.value) {
    await deletePermanent(taskToDeletePermanent.value)
    showDeletePermanentModal.value = false
    taskToDeletePermanent.value = null
  }
}

const goToPage = (page: number) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const changeFilter = (newFilter: typeof filterType.value) => {
  filterType.value = newFilter
  currentPage.value = 1
}

onMounted(() => fetchTasks())
</script>

<template>
  <main class="container">
    <div class="header">
      <h1>📋 Mis Tareas</h1>
      <p class="subtitle">Organiza tu día y mantente productivo</p>
      <p class="task-count">{{ pendingCount }} tarea{{ pendingCount !== 1 ? 's' : '' }} pendiente{{ pendingCount !== 1 ? 's' : '' }}</p>
    </div>

    <TaskForm @submit="createTask" />

    <div class="filters">
      <button 
        :class="['filter-btn', { active: filterType === 'all' }]"
        @click="changeFilter('all')"
      >
        Todas <span class="badge">{{ tasks.filter(t => !t.isDeleted).length }}</span>
      </button>
      <button 
        :class="['filter-btn', { active: filterType === 'active' }]"
        @click="changeFilter('active')"
      > 
        Activas <span class="badge">{{ tasks.filter(t => !t.isDeleted && !t.completed && t.active).length }}</span>
      </button>
      <button 
        :class="['filter-btn', { active: filterType === 'completed' }]"
        @click="changeFilter('completed')"
      >
        Completadas <span class="badge">{{ completedCount }}</span>
      </button>
      <button 
        :class="['filter-btn', { active: filterType === 'inactive' }]"
        @click="changeFilter('inactive')"
      >
        Inactivas <span class="badge">{{ inactiveCount }}</span>
      </button>
      <button 
        :class="['filter-btn', { active: filterType === 'deleted' }]"
        @click="changeFilter('deleted')"
      >
        Eliminadas <span class="badge">{{ deletedCount }}</span>
      </button>
    </div>

    <p v-if="loading" class="loading">Cargando...</p>
    <p v-if="error" class="error">{{ error }}</p>

    <TaskList
      v-if="!loading && paginatedTasks.length > 0"
      :tasks="paginatedTasks"
      :is-deleted-view="filterType === 'deleted'"
      @complete="completeTask"
      @remove="confirmDelete"
      @recover="handleRecover"
      @remove-permanent="handleRemovePermanent"
      @toggle-active="handleToggleActive"
    />
    
    <!-- Paginación -->
    <div v-if="!loading && totalPages > 1" class="pagination">
      <button 
        class="pagination-btn"
        :disabled="currentPage === 1"
        @click="goToPage(currentPage - 1)"
      >
        ← Anterior
      </button>
      
      <div class="pagination-info">
        Página {{ currentPage }} de {{ totalPages }}
      </div>
      
      <button 
        class="pagination-btn"
        :disabled="currentPage === totalPages"
        @click="goToPage(currentPage + 1)"
      >
        Siguiente →
      </button>
    </div>
    
    <div v-else-if="!loading" class="empty-state">
      <p>{{ filterType === 'deleted' ? 'No hay tareas eliminadas' : 'No hay tareas para mostrar' }}</p>
    </div>

    <!-- Modal de confirmación -->
    <div v-if="showDeleteModal" class="modal-overlay" @click.self="showDeleteModal = false">
      <div class="modal">
        <h2>¿Eliminar tarea?</h2>
        <p>Esta acción no se puede deshacer inmediatamente, pero podrás recuperarla desde la sección de Eliminadas.</p>
        <div class="modal-actions">
          <button class="btn-cancel" @click="showDeleteModal = false">Cancelar</button>
          <button class="btn-delete-confirm" @click="handleDelete">Eliminar</button>
        </div>
      </div>
    </div>

    <!-- Modal de confirmación - Delete Permanente -->
    <div v-if="showDeletePermanentModal" class="modal-overlay" @click.self="showDeletePermanentModal = false">
      <div class="modal modal-danger">
        <h2>⚠️ Eliminar Permanentemente</h2>
        <p>Esta acción no se puede deshacer. La tarea será eliminada por completo y no podrá ser recuperada.</p>
        <div class="modal-actions">
          <button class="btn-cancel" @click="showDeletePermanentModal = false">Cancelar</button>
          <button class="btn-delete-permanent-confirm" @click="handleDeletePermanent">Eliminar Definitivamente</button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f0f4f8 0%, #d9e2ec 100%);
  padding: 40px 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.header {
  text-align: center;
  margin-bottom: 30px;
  width: 100%;
  max-width: 1100px;
}

h1 {
  font-size: 48px;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 10px 0;
}

.subtitle {
  font-size: 18px;
  color: #64748b;
  margin: 0 0 10px 0;
}

.task-count {
  font-size: 14px;
  color: #94a3b8;
  margin: 0;
}

.filters {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin: 30px 0;
  flex-wrap: wrap;
  width: 100%;
  max-width: 1100px;
}

.filter-btn {
  padding: 12px 24px;
  border: none;
  border-radius: 24px;
  background: white;
  color: #64748b;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  gap: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.filter-btn.active {
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  color: white;
  box-shadow: 0 4px 8px rgba(59, 130, 246, 0.3);
}

.filter-btn:hover:not(.active) {
  background: #f8fafc;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
}

.badge {
  background: rgba(0, 0, 0, 0.1);
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

.filter-btn.active .badge {
  background: rgba(255, 255, 255, 0.3);
}

.loading, .error {
  text-align: center;
  padding: 20px;
  width: 100%;
  max-width: 1100px;
}

.error {
  color: #dc2626;
}

.empty-state {
  text-align: center;
  padding: 40px 20px;
  color: #94a3b8;
  width: 100%;
  max-width: 1100px;
}

/* Modal */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: white;
  border-radius: 12px;
  padding: 32px;
  max-width: 400px;
  box-shadow: 0 20px 25px rgba(0, 0, 0, 0.15);
}

.modal h2 {
  margin: 0 0 12px 0;
  color: #1e293b;
  font-size: 20px;
}

.modal p {
  color: #64748b;
  margin: 0 0 24px 0;
  line-height: 1.6;
}

.modal-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.btn-cancel {
  padding: 10px 20px;
  border: 1px solid #e2e8f0;
  background: white;
  color: #64748b;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s;
}

.btn-cancel:hover {
  background: #f1f5f9;
}

.btn-delete-confirm {
  padding: 10px 20px;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s;
}

.btn-delete-confirm:hover {
  background: #dc2626;
}

.modal-danger {
  border: 2px solid #fecaca;
}

.modal-danger h2 {
  color: #dc2626;
}

.btn-delete-permanent-confirm {
  padding: 10px 20px;
  background: #dc2626;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.3s;
}

.btn-delete-permanent-confirm:hover {
  background: #b91c1c;
}

/* Paginación */
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
  margin: 30px 0;
  width: 100%;
  max-width: 1100px;
}

.pagination-btn {
  padding: 10px 20px;
  border: 2px solid #3b82f6;
  background: white;
  color: #3b82f6;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s;
  font-size: 14px;
}

.pagination-btn:hover:not(:disabled) {
  background: #3b82f6;
  color: white;
  box-shadow: 0 4px 8px rgba(59, 130, 246, 0.3);
}

.pagination-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.pagination-info {
  color: #64748b;
  font-size: 14px;
  font-weight: 500;
  min-width: 140px;
  text-align: center;
}
</style>