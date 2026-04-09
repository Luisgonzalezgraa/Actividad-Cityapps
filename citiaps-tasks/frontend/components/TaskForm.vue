<script setup lang="ts">
import { ref } from 'vue'

const emit = defineEmits<{
  submit: [payload: { title: string; description: string; tags: string[] }]
}>()

const title = ref('')
const tagsInput = ref('')

const onSubmit = () => {
  if (!title.value.trim()) return

  const tags = tagsInput.value
    .split(',')
    .map(tag => tag.trim())
    .filter(Boolean)

  emit('submit', {
    title: title.value,
    description: '',
    tags
  })

  title.value = ''
  tagsInput.value = ''
}
</script>

<template>
  <form class="task-form" @submit.prevent="onSubmit">
    <div class="form-main">
      <input 
        v-model="title" 
        type="text" 
        class="task-input"
        placeholder="¿Qué necesitas hacer hoy?" 
      />
      <button type="submit" class="btn-add">
        <span>+</span> Agregar
      </button>
    </div>
    <input 
      v-model="tagsInput" 
      type="text" 
      class="tags-input"
      placeholder="Agregar etiquetas (presiona Enter)..." 
    />
  </form>
</template>

<style scoped>
.task-form {
  width: 100%;
  max-width: 1060px;
  margin: 0 auto 30px;
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.form-main {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
}

.task-input {
  flex: 1;
  min-width: 0;
  padding: 12px 16px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 16px;
  transition: all 0.3s;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  color: #1e293b;
  box-sizing: border-box;
}

.task-input:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.task-input::placeholder {
  color: #cbd5e1;
}

.btn-add {
  padding: 12px 24px;
  background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
  align-items: center;
  gap: 8px;
  white-space: nowrap;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

.btn-add:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(59, 130, 246, 0.4);
}

.btn-add:active {
  transform: translateY(0);
}

.btn-add span {
  font-size: 20px;
  font-weight: 700;
}

.tags-input {
  width: 100%;
  box-sizing: border-box;
  padding: 12px 16px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.3s;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  color: #64748b;
}

.tags-input:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  color: #1e293b;
}

.tags-input::placeholder {
  color: #cbd5e1;
}

@media (max-width: 640px) {
  .task-form {
    padding: 16px;
  }

  .form-main {
    flex-direction: column;
  }

  .btn-add {
    width: 100%;
    justify-content: center;
  }
}
</style>