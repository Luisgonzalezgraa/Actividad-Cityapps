<script setup lang="ts">
import type { Task } from '~/composables/useTasks'
import RoundCheckbox from '~/components/RoundCheckbox.vue'

defineProps<{
  tasks: Task[]
  isDeletedView?: boolean
}>()

const emit = defineEmits<{
  complete: [id: string]
  remove: [id: string]
  removePermanent: [id: string]
  recover: [id: string]
  toggleActive: [id: string]
}>()

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  const day = String(date.getDate()).padStart(2, '0')
  const months = ['ene', 'feb', 'mar', 'abr', 'may', 'jun', 'jul', 'ago', 'sep', 'oct', 'nov', 'dic']
  const month = months[date.getMonth()]
  const year = date.getFullYear()
  return `${day} ${month} ${year}`
}

const getStatusColor = (completed: boolean, active: boolean) => {
  if (!active) return '#94a3b8'
  return completed ? '#10b981' : '#3b82f6'
}

const getStatusText = (completed: boolean, active: boolean) => {
  if (!active) return 'Inactiva'
  return completed ? 'Completada' : 'Activa'
}
</script>

<template>
  <div class="task-list">
    <div v-for="task in tasks" :key="task._id" class="task-item" :class="{ deleted: isDeletedView, inactive: !isDeletedView && !task.active }">
      <div v-if="!isDeletedView" class="task-checkbox">
        <RoundCheckbox 
          :model-value="task.completed"
          @update:model-value="emit('complete', task._id)"
        />
      </div>
      
      <div class="task-content">
        <h3 :class="{ completed: task.completed }">{{ task.title }}</h3>
        <div class="task-meta">
          <span class="date">📅 {{ formatDate(task.createdAt) }}</span>
          <span v-if="!isDeletedView"
            class="status"
            :style="{ '--status-color': getStatusColor(task.completed, task.active) }"
          >
            ⏱ {{ getStatusText(task.completed, task.active) }}
          </span>
        </div>
        <div v-if="task.tags && task.tags.length > 0" class="tags">
          <span v-for="tag in task.tags" :key="tag" class="tag">{{ tag }}</span>
        </div>
      </div>

      <div v-if="isDeletedView" class="actions-deleted">
        <button class="btn-recover" @click="emit('recover', task._id)" title="Recuperar tarea">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path d="M12 2v10"></path>
            <path d="M5.5 5.5a7 7 0 1 0 13 0"></path>
          </svg>
        </button>
        <button class="btn-delete-permanent" @click="emit('removePermanent', task._id)" title="Eliminar permanentemente">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path d="M3 6h18"></path>
            <path d="M8 6V4h8v2"></path>
            <path d="M6 6l1 14h10l1-14"></path>
          </svg>
        </button>
      </div>
      <div v-else class="actions-active">
        <button class="btn-toggle-active" :class="{ inactive: !task.active }" @click="emit('toggleActive', task._id)" :title="task.active ? 'Desactivar' : 'Activar'">
          <svg viewBox="0 0 24 24" class="icon">
            <path d="M12 2v8" />
            <path d="M7 5a7 7 0 1 0 10 0" />
          </svg>
        </button>
        <button class="btn-delete" @click="emit('remove', task._id)">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path d="M3 6h18"></path>
            <path d="M8 6V4h8v2"></path>
            <path d="M6 6l1 14h10l1-14"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.task-list {
  max-width: 1060px;
  margin: 0 auto;
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.task-item {
  background: white;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  transition: all 0.3s;
  border: 2px solid transparent;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.task-item.inactive {
  background: #f8fafc;
  opacity: 0.7;
}

.task-item.deleted {
  border: 2px solid #e0e7ff;
  background: #f8fafc;
}

.task-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.task-checkbox {
  flex-shrink: 0;
  padding-top: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.task-content {
  flex: 1;
}

h3 {
  margin: 0 0 12px 0;
  font-size: 16px;
  font-weight: 500;
  color: #1e293b;
  transition: all 0.3s;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

h3.completed {
  text-decoration: line-through;
  color: #94a3b8;
}

.task-meta {
  display: flex;
  gap: 16px;
  align-items: center;
  flex-wrap: wrap;
  margin-bottom: 8px;
}

.date {
  font-size: 13px;
  color: #64748b;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.status {
  font-size: 13px;
  color: var(--status-color);
  font-weight: 500;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-top: 8px;
}

.tag {
  background: #f1f5f9;
  color: #64748b;
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
  border: 1px solid #e2e8f0;
}

.btn-delete {
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border: none;
  background: #fce7e6;
  color: #ff6b6b;
  border-radius: 50%;
  cursor: pointer;
  font-size: 18px;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  visibility: hidden;
}

.task-item:hover .btn-delete {
  opacity: 1;
  visibility: visible;
}

.btn-delete:hover {
  background: #ffcccb;
  transform: scale(1.05);
}

.btn-delete:active {
  transform: scale(0.95);
}

.actions-active {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-left: auto;
}

.btn-toggle-active {
  width: 50px;
  height: 50px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  background: #d1fae5;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08);
  cursor: pointer;
  transition: all 0.3s;
  flex-shrink: 0;
  opacity: 0;
  visibility: hidden;
}

.task-item:hover .btn-toggle-active {
  opacity: 1;
  visibility: visible;
}

.btn-toggle-active:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
}

.btn-toggle-active.inactive {
  background: #f1f5f9;
}

.btn-toggle-active.inactive:hover {
  transform: scale(1.05);
}

.btn-toggle-active:active {
  transform: scale(0.95);
}

.icon {
  width: 24px;
  height: 24px;
  stroke: #22c55e;
  stroke-width: 2.5;
  fill: none;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.btn-toggle-active.inactive .icon {
  stroke: #94a3b8;
}

.task-item.deleted {
  border: 2px solid #e0e7ff;
  background: #f8fafc;
}

.actions-deleted {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-left: auto;
}

.btn-recover {
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border: none;
  background: #e0e7ff;
  color: #6366f1;
  border-radius: 50%;
  cursor: pointer;
  font-size: 18px;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  visibility: hidden;
}

.task-item:hover .btn-recover {
  opacity: 1;
  visibility: visible;
}

.btn-recover:hover {
  background: #c7d2fe;
  transform: scale(1.05);
}

.btn-recover:active {
  transform: scale(0.95);
}

.btn-delete-permanent {
  flex-shrink: 0;
  width: 40px;
  height: 40px;
  border: none;
  background: #fce7e6;
  color: #ff6b6b;
  border-radius: 50%;
  cursor: pointer;
  font-size: 18px;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  visibility: hidden;
}

.task-item:hover .btn-delete-permanent {
  opacity: 1;
  visibility: visible;
}

.btn-delete-permanent:hover {
  background: #ffcccb;
  transform: scale(1.05);
}

.btn-delete-permanent:active {
  transform: scale(0.95);
}

svg {
  display: block;
  width: 20px;
  height: 20px;
  stroke-linecap: round;
  stroke-linejoin: round;
}

@media (max-width: 640px) {
  .task-item {
    padding: 16px;
    gap: 12px;
  }

  h3 {
    font-size: 15px;
  }

  .task-meta {
    gap: 12px;
  }

  .date, .status {
    font-size: 12px;
  }
}
</style>