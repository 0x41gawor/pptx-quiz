<template>
  <div class="admin-page">
    <AdminNav />

    <div class="admin-content">
      <h1>Completions</h1>

      <div class="toolbar">
        <input
          v-model="query"
          type="text"
          placeholder="Szukaj (imię + nazwisko)"
          @input="load"
        />
      </div>

      <table class="completions-table">
        <thead>
          <tr>
            <th>Id</th>
            <th>Imię</th>
            <th>Nazwisko</th>
            <th>Telefon</th>
            <th>Data</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="c in completions" :key="c.id">
            <td>{{ c.id }}</td>
            <td>{{ c.firstname }}</td>
            <td>{{ c.lastname }}</td>
            <td>{{ c.phone_number }}</td>
            <td>{{ c.date }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import AdminNav from '@/components/layout/AdminNav.vue'
import { ref, onMounted } from 'vue'
import type { Completion } from '@/models/completion'
import { listCompletions } from '@/api/completionsApi'

const completions = ref<Completion[]>([])
const query = ref('')

async function load() {
  completions.value = await listCompletions(query.value || undefined)
}

onMounted(load)
</script>

<style scoped>
.admin-page {
  min-height: 100vh;
  background: var(--color-admin-bg);
}

.admin-content {
  max-width: 960px;
  margin: 2rem auto;
  padding: 1.5rem;
  background: var(--color-admin-card);
  border-radius: 1rem;
  border: 1px solid var(--color-admin-border);
}

.toolbar {
  margin: 1rem 0;
}

.completions-table {
  width: 100%;
  border-collapse: collapse;
}

.completions-table th,
.completions-table td {
  border: 1px solid var(--color-admin-border);
  padding: 0.5rem;
  text-align: left;
}
</style>
