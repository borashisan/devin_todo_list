<template>
  <div class="min-h-screen bg-gradient-to-br from-indigo-900 via-purple-900 to-pink-800">
    <div class="container mx-auto px-4 py-12">
      <!-- Header -->
      <div class="text-center mb-10">
        <h1 class="text-4xl font-bold text-white mb-2">📝 Todo App</h1>
        <p class="text-purple-200">Manage your tasks efficiently</p>
      </div>

      <!-- Todo Card -->
      <div class="max-w-2xl mx-auto">
        <div class="bg-white/10 backdrop-blur-lg rounded-2xl shadow-2xl border border-white/20 overflow-hidden">
          <!-- Input Form -->
          <div class="p-6 border-b border-white/10">
            <div class="flex gap-3">
              <input
                v-model="newTodoTitle"
                type="text"
                placeholder="Add a new task..."
                class="flex-1 px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-purple-300 focus:outline-none focus:ring-2 focus:ring-purple-400 focus:border-transparent transition-all"
                @keyup.enter="addTodo"
              />
              <button
                @click="addTodo"
                class="px-6 py-3 bg-gradient-to-r from-purple-500 to-pink-500 text-white font-semibold rounded-xl hover:from-purple-600 hover:to-pink-600 transform hover:scale-105 transition-all duration-200 shadow-lg"
              >
                追加
              </button>
            </div>
          </div>

          <!-- Todo List -->
          <div class="p-6">
            <ul class="space-y-3">
              <li
                v-for="todo in todos"
                :key="todo.id"
                class="group flex items-center gap-4 p-4 bg-white/5 hover:bg-white/10 rounded-xl border border-white/10 transition-all duration-200"
              >
                <!-- Checkbox -->
                <button
                  @click="toggleTodo(todo.id)"
                  class="flex-shrink-0 w-6 h-6 rounded-full border-2 flex items-center justify-center transition-all duration-200"
                  :class="todo.is_completed
                    ? 'bg-gradient-to-r from-green-400 to-emerald-500 border-transparent'
                    : 'border-purple-400 hover:border-purple-300'"
                >
                  <svg
                    v-if="todo.is_completed"
                    class="w-4 h-4 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
                  </svg>
                </button>

                <!-- Title -->
                <span
                  class="flex-1 text-lg transition-all duration-200"
                  :class="todo.is_completed ? 'text-purple-300 line-through' : 'text-white'"
                >
                  {{ todo.title }}
                </span>

                <!-- Delete Button -->
                <button
                  @click="deleteTodo(todo.id)"
                  class="flex-shrink-0 px-3 py-1.5 text-sm text-red-300 hover:text-white hover:bg-red-500/50 rounded-lg opacity-0 group-hover:opacity-100 transition-all duration-200"
                >
                  削除
                </button>
              </li>

              <!-- Empty State -->
              <li v-if="todos.length === 0" class="text-center py-12">
                <p class="text-purple-300 text-lg">No tasks yet. Add one above!</p>
              </li>
            </ul>
          </div>

          <!-- Footer Stats -->
          <div class="px-6 py-4 bg-white/5 border-t border-white/10">
            <div class="flex justify-between text-sm text-purple-300">
              <span>{{ completedCount }} / {{ todos.length }} completed</span>
              <span v-if="remainingCount > 0">{{ remainingCount }} remaining</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Todo {
  id: string
  title: string
  is_completed: boolean
}

// Mock data (will be replaced with API calls later)
const todos = ref<Todo[]>([
  { id: '1', title: 'Buy groceries', is_completed: false },
  { id: '2', title: 'Read a book', is_completed: true },
  { id: '3', title: 'Write code', is_completed: false },
])

const newTodoTitle = ref('')

const completedCount = computed(() => todos.value.filter(t => t.is_completed).length)
const remainingCount = computed(() => todos.value.filter(t => !t.is_completed).length)

function addTodo() {
  if (!newTodoTitle.value.trim()) return

  const newTodo: Todo = {
    id: Date.now().toString(),
    title: newTodoTitle.value.trim(),
    is_completed: false,
  }
  todos.value.push(newTodo)
  newTodoTitle.value = ''
}

function toggleTodo(id: string) {
  const todo = todos.value.find(t => t.id === id)
  if (todo) {
    todo.is_completed = !todo.is_completed
  }
}

function deleteTodo(id: string) {
  todos.value = todos.value.filter(t => t.id !== id)
}
</script>
