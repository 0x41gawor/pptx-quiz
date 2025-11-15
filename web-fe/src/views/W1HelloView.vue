<template>
  <div class="hello-view">
    <div class="content">
      <h1 class="title">{{ t('hello_title').toUpperCase() }}</h1>

      <p class="subtitle">
        {{ t('hello_subtitle').toUpperCase() }}
      </p>

      <div class="flags">
        <div class="flag-wrapper" @click="select('pl')">
          <img src="/flags/pl.png" alt="Polish flag" />
        </div>

        <div class="flag-wrapper" @click="select('en')">
          <img src="/flags/en.png" alt="UK flag" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Language } from '@/models/common'
import { useI18n } from '@/i18n/lang'
import { useRouter } from 'vue-router'
import { quizSession } from '@/stores/quizSession'

const router = useRouter()
const { setLanguage, t } = useI18n()

function select(lang: Language) {
  setLanguage(lang)
  quizSession.language = lang
  router.push({ name: 'user-data' })
}
</script>

<style scoped>
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
  padding-left: 5vw; /* lekki padding jak w mocku */
}

/* wyłączamy centralne grupowanie — w mocku każdy element jest inaczej położony */
.content {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 60px;
}

/* --- TYTUŁ --- */
.title {
  font-size: 92px;
  font-weight: 400;
  letter-spacing: 2px;
  text-align: right;

  /* gradient tekstu */
  background: linear-gradient(
    to right,
    var(--color-pink-light),
    var(--color-orange)
  );
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;

  /* pozycjonowanie 1:1 jak na mocku */
  width: fit-content;
  margin-left: 60vw; /* tytuł najmocniej przesunięty w prawo */
  margin-right: 5vw;
}

/* --- PODTYTUŁ --- */
.subtitle {
  font-size: 32px;
  color: white;
  letter-spacing: 1px;
  width: fit-content;

  /* mock: subtitle jest bardziej w lewo niż tytuł */
  margin-left: 48vw;
  margin-top: 20px;
}

/* --- FLAGI --- */
.flags {
  display: flex;
  gap: 60px;

  /* mock: flagi są jeszcze bardziej w lewo */
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