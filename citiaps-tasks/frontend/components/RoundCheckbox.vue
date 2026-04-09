<script setup lang="ts">
const model = defineModel<boolean>({ default: false })
</script>

<template>
  <label class="round-check">
    <input v-model="model" type="checkbox" />
    <span class="check-ui">
      <svg viewBox="0 0 24 24" class="tick" aria-hidden="true">
        <path d="M5 12.5l4.2 4.2L19 7.8" />
      </svg>
    </span>
  </label>
</template>

<style scoped>
.round-check {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  user-select: none;
}

.round-check input {
  display: none;
}

.check-ui {
  width: 25px;
  height: 25px;
  border-radius: 999px;
  border: 2px solid #3b82f6;
  background: #eff6ff;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  position: relative;
  box-shadow: 0 6px 14px rgba(0, 0, 0, 0.08);
  transition:
    transform 0.18s ease,
    background 0.22s ease,
    box-shadow 0.22s ease,
    border-color 0.22s ease;
}

.tick {
  width: 18px;
  height: 18px;
  opacity: 0;
  transform: scale(0.4) rotate(-18deg);
  transition:
    opacity 0.16s ease,
    transform 0.26s cubic-bezier(.2, 1.4, .4, 1);
}

.tick path {
  fill: none;
  stroke: #2563eb;
  stroke-width: 3.2;
  stroke-linecap: round;
  stroke-linejoin: round;
  stroke-dasharray: 22;
  stroke-dashoffset: 22;
  transition: stroke-dashoffset 0.28s ease 0.04s;
}

/* al presionar */
.round-check:active .check-ui {
  transform: scale(0.9);
}

/* checked */
.round-check input:checked + .check-ui {
  background: #dbeafe;
  border-color: #2563eb;
  animation: pop 0.28s ease;
}

.round-check input:checked + .check-ui .tick {
  opacity: 1;
  transform: scale(1) rotate(0deg);
}

.round-check input:checked + .check-ui .tick path {
  stroke-dashoffset: 0;
}

/* foco teclado */
.round-check input:focus-visible + .check-ui {
  outline: 3px solid rgba(59, 130, 246, 0.22);
  outline-offset: 3px;
}

@keyframes pop {
  0% {
    transform: scale(0.82);
  }
  55% {
    transform: scale(1.08);
  }
  100% {
    transform: scale(1);
  }
}
</style>
