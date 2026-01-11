<template>
  <div class="admin-page">
    <AdminNav />

    <div class="admin-content">
      <h1>Plik z pytaniami</h1>
      <div class="actions">
        <button @click="handleDownload">Podgląd pliku</button>

        <label class="upload-label">
          Wgraj plik
          <input type="file" accept=".csv" @change="handleFileChange" hidden />
        </label>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import AdminNav from '@/components/layout/AdminNav.vue'
import { downloadQuizFile, uploadQuizFile } from '@/api/quizApi'

async function handleDownload() {
  try {
    const blob = await downloadQuizFile()
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'quiz.csv'
    a.click()
    URL.revokeObjectURL(url)
  } catch (err) {
    console.error(err)
    alert('Nie udało się pobrać pliku')
  }
}

async function handleFileChange(event: Event) {
  const target = event.target as HTMLInputElement
  if (!target.files?.length) return

  const file = target.files[0]

  try {
    // TODO: walidacja pliku po stronie FE
    if (!file) return;
    await uploadQuizFile(file)
    alert('Plik został wgrany')
  } catch (err) {
    console.error(err)
    alert('Nie udało się wgrać pliku')
  } finally {
    target.value = ''
  }
}
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

.actions {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

button,
.upload-label {
  border-radius: 0.5rem;
  border: 1px solid var(--color-admin-border);
  background: #fff;
  padding: 0.5rem 1rem;
}
.upload-label {
  cursor: pointer;
}
</style>
