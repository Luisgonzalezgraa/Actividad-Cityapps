# CITIAPS - Task Management Application

Una aplicación web moderna para la gestión completa de tareas, con soporte para crear, completar, archivar, recuperar y eliminar tareas de forma permanente.

## Características

-  Crear, editar y eliminar tareas
-  Marcar tareas como completadas
-  Filtrar tareas (Todas, Activas, Completadas, Inactivas, Eliminadas)
-  Paginación de 5 tareas por página
-  Soft delete (tareas recuperables)
-  Hard delete (eliminación permanente)
-  Sistema de etiquetas
-  Interfaz moderna y responsiva
-  Base de datos MongoDB con persistencia

##  Stack Tecnológico

**Backend:**
- Go 1.25
- Gin (framework web)
- MongoDB Driver
- CORS habilitado

**Frontend:**
- Nuxt 3
- Vue 3
- TypeScript
- CSS personalizado (sin frameworks)
- Font: Inter (Google Fonts)

**Infraestructura:**
- Docker & Docker Compose
- MongoDB 7

---

##  Requisitos Previos

- **Docker & Docker Compose** (recomendado para ejecución rápida)
- O instalados localmente:
  - Go 1.25+
  - Node.js 20+
  - MongoDB 7+
  - npm o yarn

---

##  Guía de Ejecución

###  Opción 1: Con Docker (Recomendado - Más Rápido!)

**Levanta toda la aplicación con un solo comando:**

```bash
cd /ruta/a/citiaps-tasks
docker compose up --build
```

**Servicios que se inician automáticamente:**
-  MongoDB en `localhost:27017`
-  Backend en `http://localhost:8080`
-  Frontend en `http://localhost:3000`

**Verificar que todo está funcionando:**
```bash
# En otra terminal
curl http://localhost:8080/health
```

**Para detener:**
```bash
docker compose down
```

**Para detener y eliminar volúmenes (reset completo):**
```bash
docker compose down -v
```

---

### Opción 2: Ejecución Manual (Local)

#### Configurar MongoDB

**Opción A: Con Docker (solo MongoDB)**
```bash
docker run -d \
  --name tasks-mongo \
  -p 27017:27017 \
  -v mongo_data:/data/db \
  mongo:7
```

**Opción B: Instalado localmente**

*Windows:*
```bash
mongod
```

*macOS (con Homebrew):*
```bash
brew services start mongodb-community
```

*Linux:*
```bash
sudo systemctl start mongod
```

**Verificar conexión:**
```bash
mongosh  # o mongo en versiones antiguas
```

---

####  Cómo Ejecutar el Backend

```bash
cd backend

# 1. Copiar archivo de configuración
cp .env.example .env

# 2. Configurar variables de entorno en .env:
# PORT=8080
# MONGO_URI=mongodb://localhost:27017
# MONGO_DB=tasksdb
# MONGO_COLLECTION=tasks

# 3. Descargar dependencias
go mod download

# 4. Ejecutar servidor en modo desarrollo
go run cmd/server/main.go
```

**Expected output:**
```
[GIN-debug] Loaded HTML Templates (0): 
[GIN-debug] Listening and serving HTTP on :8080
```

---

####  Cómo Ejecutar el Frontend

```bash
cd frontend

# 1. Instalar dependencias
npm install

# 2. Ejecutar en modo desarrollo con hot-reload
npm run dev -- --host 0.0.0.0
```

**Expected output:**
```
   Local:    http://localhost:3000/
   Network:  http://192.168.x.x:3000/
```

**Abre en el navegador:** `http://localhost:3000`

---

##  Estructura del Proyecto

```
citiaps-tasks/
├── backend/                           # Backend Go
│   ├── cmd/
│   │   └── server/
│   │       └── main.go               # Punto de entrada
│   ├── internal/
│   │   ├── config/                   # Configuración (variables de entorno)
│   │   ├── db/                       # Conexión a MongoDB
│   │   ├── handlers/                 # Manejadores HTTP
│   │   ├── repositories/             # Acceso a datos (CRUD en BD)
│   │   ├── routes/                   # Definición de rutas
│   │   ├── services/                 # Lógica de negocio
│   │   └── models/                   # Estructuras de datos
│   ├── Dockerfile                    # Imagen Docker optimizada
│   ├── .dockerignore                 # Archivos a ignorar en Docker
│   ├── go.mod / go.sum               # Dependencias
│   └── .env.example                  # Template de configuración
│
├── frontend/                          # Frontend Nuxt/Vue
│   ├── pages/
│   │   └── index.vue                 # Página principal
│   ├── components/
│   │   ├── TaskForm.vue              # Formulario crear tareas
│   │   ├── TaskList.vue              # Listado de tareas
│   │   └── RoundCheckbox.vue         # Checkbox personalizado
│   ├── composables/
│   │   └── useTasks.ts               # Lógica compartida (API)
│   ├── Dockerfile                    # Imagen Docker optimizada
│   ├── .dockerignore                 # Archivos a ignorar
│   ├── nuxt.config.ts                # Configuración Nuxt
│   ├── package.json                  # Dependencias npm
│   └── .env.example                  # Template de configuración
│
├── docker-compose.yml                # Orquestación de servicios
└── README.md                         # Este archivo
```

---

##  Endpoints del Backend

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| **GET** | `/health` | Verificar estado del servidor |
| **POST** | `/tasks` | Crear nueva tarea |
| **GET** | `/tasks` | Obtener todas las tareas |
| **GET** | `/tasks/:id` | Obtener tarea por ID |
| **PUT** | `/tasks/:id/complete` | Marcar como completada |
| **PUT** | `/tasks/:id/recover` | Recuperar tarea eliminada |
| **PUT** | `/tasks/:id/toggle-active` | Activar/Desactivar |
| **DELETE** | `/tasks/:id` | Eliminar (soft delete) |
| **DELETE** | `/tasks/:id/permanent` | Eliminar permanentemente |

**Ejemplo de crear tarea:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Hacer compras",
    "description": "Ir al supermercado",
    "tags": ["compras", "urgente"]
  }'
```

---

##  Sistema de Paginación

### Funcionamiento

La paginación muestra **5 tareas por página** y se aplicа **después de los filtros**. Esto garantiza que:

- Si aplicas un filtro (ej: "Inactivas") y tienes 6 tareas inactivas, verás 5 en página 1 y 1 en página 2
- Cada filtro tiene su propia paginación independiente
- Al cambiar de filtro, se resetea a la página 1

### Controles de Paginación

En la interfaz encontrarás:
- Botón **"← Anterior"** (deshabilitado en página 1)
- Indicador visual: **"Página X de Y"**
- Botón **"Siguiente →"** (deshabilitado en última página)

### Ejemplo de Uso

```
1. Tienes 12 tareas totales no eliminadas
2. Filtro "Todas" → Página 1 muestra tareas 1-5, Página 2 muestra tareas 6-10, Página 3 muestra tareas 11-12
3. Cambias a filtro "Completadas" → Se reinicia en Página 1 con tus tareas completadas paginadas
4. Tienes 6 tareas completadas → Página 1 muestra 5, Página 2 muestra 1
```

---

##  Modelo de Datos (MongoDB)

**Colección: `tasks`**

```json
{
  "_id": "ObjectId",
  "title": "string",
  "description": "string",
  "completed": "boolean",
  "active": "boolean",
  "isDeleted": "boolean",
  "tags": ["string"],
  "createdAt": "ISODate",
  "updatedAt": "ISODate"
}
```


##  Flujo de Ejecución (Crear Tarea)

```
1. Usuario escribe en TaskForm.vue
                 ↓
2. Presiona botón "+ Agregar"
                 ↓
3. @submit emite evento con datos
                 ↓
4. index.vue recibe en @submit="createTask"
                 ↓
5. useTasks.ts hace POST HTTP /tasks
                 ↓
6. Backend:
   - Handler recibe JSON
   - Service valida datos
   - Repository inserta en MongoDB
                 ↓
7. MongoDB retorna documento creado
                 ↓
8. Backend retorna tarea (HTTP 201)
                 ↓
9. Frontend ejecuta fetchTasks()
                 ↓
10. TaskList.vue muestra nueva tarea 
```

---

##  Tiempo Aproximado de Desarrollo

| Fase | Tiempo |
|------|--------|
| Setup y scaffolding | 30 min |
| Backend (CRUD + MongoDB) | 90 min |
| Frontend (componentes base) | 60 min |
| Soft/Hard delete + modales | 45 min |
| Styling y diseño UI/UX | 75 min |
| Testing y ajustes | 60 min |
| Dockerización y optimización | 40 min |
| **TOTAL** | **~12 horas** |

---

##  Troubleshooting

###  Error: "Connection refused" en MongoDB

**Solución:**
```bash
# Verificar MongoDB está corriendo
mongosh

# Si no, iniciar con Docker:
docker run -d -p 27017:27017 mongo:7
```

###  Error: "Port already in use"

**Solución:**
```bash
# Cambiar puerto en .env o ejecutar en otro:
npm run dev -- --port 3001
```

###  Frontend no conecta con backend

**Verificar:**
1. Backend corriendo: `curl http://localhost:8080/health`
2. Variable de entorno: `NUXT_PUBLIC_API_BASE=http://localhost:8080`
3. CORS habilitado en backend

##  Build de Producción

**Con Docker:**
```bash
docker compose build
docker compose up -d
```

**Variables de entorno necesarias:**
```bash
MONGO_URI=mongodb://prod-mongo-host:27017
NUXT_PUBLIC_API_BASE=https://api.tudominio.com
PORT=8080
```

---

##  Referencias Útiles

- [Go & Gin Framework](https://gin-gonic.com/)
- [Nuxt 3 Documentation](https://nuxt.com/)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Vue 3 Reactive API](https://vuejs.org/)

---
