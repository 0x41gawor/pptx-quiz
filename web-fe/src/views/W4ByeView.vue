<template>
  <div class="bye-view" @click="goHome">
    <div class="bye-content">
      <h1 class="bye-title">
        THANK YOU
      </h1>

      <p class="bye-subtitle">
        PLEASE GO TO THE<br />
        RECEPTION FOR FURTHER<br />
        INSTRUCTIONS
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'

const router = useRouter()

function goHome() {
  router.push({ name: 'hello' })
}

onMounted(() => {
  setTimeout(goHome, 10_000)
})
</script>

<style scoped>
.bye-view {
  min-height: 100vh;
  width: 100%;
  background: var(--color-dark-bg);

  display: flex;
  align-items: center;
  justify-content: flex-start;

  padding-left: 10vw; /* przesuniecie jak w mocku */
}

.bye-content {
  display: flex;
  flex-direction: column;
  gap: 64px;
}

.bye-title {
  font-size: 96px;
  font-weight: 600;
  color: #ffffff;
  letter-spacing: 2px;
  margin: 0;
}

.bye-subtitle {
  font-size: 48px;
  font-weight: 300;
  color: #ffffff;
  line-height: 1.35;
  margin: 0;
}
</style>