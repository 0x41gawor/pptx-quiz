<template>
  <div class="video-page">
    <!-- Language selection -->
    <div v-if="!selectedLang" class="hello-view">
    <div class="content">
        <h1 class="title">{{ t('video_title').toUpperCase() }}</h1>

        <p class="subtitle">
        {{ t('hello_subtitle').toUpperCase() }}
        </p>

        <div class="flags">
        <div class="flag-wrapper" @click="selectLang('pl')">
            <img src="/flags/pl.png" alt="Polish flag" />
        </div>

        <div class="flag-wrapper" @click="selectLang('en')">
            <img src="/flags/en.png" alt="UK flag" />
        </div>
        </div>
    </div>
    </div>

    <!-- Video player -->
    <div v-else class="player-wrapper">
      <video
        ref="videoEl"
        class="video-player"
        controls
        preload="metadata"
        playsinline
        @ended="onVideoEnded"
      >
        <source :src="videoSrc" type="video/mp4" />
        Twoja przeglądarka nie obsługuje wideo.
      </video>

      <button class="exit-btn" @click="exitVideo">
        Zmień język
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { Language } from '@/models/common'
import { useI18n } from '@/i18n/lang'
import { useRouter } from 'vue-router'

const router = useRouter()


const { setLanguage, t } = useI18n()

const selectedLang = ref<Language | null>(null)
const videoEl = ref<HTMLVideoElement | null>(null)

const videoSrc = computed(() => {
  if (!selectedLang.value) return ''
  return `/media/videos/intro_${selectedLang.value}.mp4`
})

function selectLang(lang: Language) {
  setLanguage(lang)
  selectedLang.value = lang
}

function exitVideo() {
  if (videoEl.value) {
    videoEl.value.pause()
    videoEl.value.currentTime = 0
  }
  selectedLang.value = null
}

function onVideoEnded() {
  router.push({ name: 'video-ended' })
}
</script>


<style scoped>
.video-page {
  width: 100vw;
  height: 100vh;
  background: var(--color-dark-bg);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* LANGUAGE SELECT */

.lang-select {
  text-align: center;
}

.lang-select h1 {
  margin-bottom: 2rem;
}

.flags {
  display: flex;
  gap: 2rem;
  justify-content: center;
}

.flag-btn {
  background: none;
  border: none;
  padding: 0;
}

.flag-btn img {
  width: 120px;
  height: auto;
}

/* VIDEO PLAYER */

.player-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
}

.video-player {
  width: 100%;
  height: 100%;
  object-fit: contain;
  background: black;
}

.exit-btn {
  position: absolute;
  top: 1rem;
  left: 1rem;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  border: none;
  padding: 0.75rem 1.25rem;
  font-size: 1rem;
  border-radius: 6px;
}

.hello-view {
  min-height: 100vh;
  width: 100%;

  background: url('/backgrounds/1.png');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;

  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding-left: 5vw;
}

.content {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 60px;
}

.title {
  font-size: 92px;
  font-weight: 400;
  letter-spacing: 2px;
  text-align: right;

  background: linear-gradient(
    to right,
    var(--color-pink-light),
    var(--color-orange)
  );
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;

  width: fit-content;
  margin-left: 60vw;
  margin-right: 5vw;
}

.subtitle {
  font-size: 32px;
  color: white;
  letter-spacing: 1px;
  width: fit-content;

  margin-left: 48vw;
  margin-top: 20px;
}

.flags {
  display: flex;
  gap: 60px;
  margin-left: 40vw;
  margin-top: 40px;
}

.flag-wrapper {
  width: 320px;
  height: 160px;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.flag-wrapper:hover {
  transform: scale(1.04);
}

.flag-wrapper img {
  width: 100%;
  height: 100%;
  border-radius: 6px;
  object-fit: cover;
}
</style>
