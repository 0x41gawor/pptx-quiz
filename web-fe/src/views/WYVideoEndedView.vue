<template>
  <div class="bye-view" @click="goHome">
    <div class="bye-content">
      <h1 class="bye-title">
        {{ t('video_ended_title') }}
      </h1>

      <p class="bye-subtitle">
        {{ t('video_ended_subtitle') }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'
import { useI18n } from '@/i18n/lang'

const router = useRouter()
const { t } = useI18n()

function goHome() {
  router.push({ name: 'video' })
}

onMounted(() => {
  setTimeout(goHome, 10_000)
})
</script>

<style scoped>
.bye-view {
  min-height: 100vh;
  width: 100%;

  background: url('/backgrounds/4.png');
  background-size: cover;
  background-position: right center;
  background-repeat: no-repeat;

  display: flex;
  align-items: center;
  justify-content: flex-start;

  padding-left: 10vw;
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