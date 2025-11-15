<template>
  <nav class="admin-nav">
    <div class="left">
      <span class="logo">pptx-quiz Admin</span>
    </div>
    <div class="right">
      <RouterLink to="/admin/file-upload">Quiz file</RouterLink>
      <RouterLink to="/admin/completions">Completions</RouterLink>
    </div>
  </nav>
</template>

<script setup lang="ts"></script>

<style scoped>
.admin-nav {
  height: 56px;
  background: #111827;
  color: #f9fafb;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 1.5rem;
}

.right {
  display: flex;
  gap: 1rem;
}

a.router-link-active {
  text-decoration: underline;
}
</style>
