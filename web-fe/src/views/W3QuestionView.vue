<template>
  <div class="question-view" v-if="currentQuestion">
    <!-- LEWA STRONA: GRADIENT + PYTANIE + PROGRESS -->
    <div class="left-panel">
      <div class="left-inner">
        <div class="top">
          <p class="counter">
            QUESTION {{ currentIndex + 1 }}/{{ totalQuestions }}
          </p>
          <h2 class="question-text">
            {{ currentQuestion.content }}
          </h2>
        </div>

        <div class="bottom">
          <div class="progress-bar">
            <div class="progress-bar-fill" :style="{ width: progressPercent + '%' }"></div>
            <span class="progress-label">
              {{ Math.round(progressPercent) }}%
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- PRAWA STRONA: ODPOWIEDZI + NEXT/REPEAT -->
    <div class="right-panel">
      <div class="answers">
<button
  v-for="(answer, index) in currentQuestion.answers"
  :key="index"
  class="answer-button"
  :class="answerButtonClass(index)"
  @click="onAnswerClick(index)"
>
  <!-- PREFIX: litera lub ikonka -->
  <span class="answer-prefix" v-if="!isSelected(index)">
    {{ prefixLetters[index] }}
  </span>

  <img
    v-if="isSelected(index) && state === 'correct'"
    class="status-img"
    src="/status/correct.png"
    alt="Correct"
  />

  <img
    v-if="isSelected(index) && state === 'wrong'"
    class="status-img"
    src="/status/wrong.png"
    alt="Wrong"
  />

  <span class="answer-text">{{ answer.content }}</span>
</button>

      </div>

      <div class="controls" v-if="state !== 'idle'">
        <button
          v-if="state === 'correct'"
          class="btn-next"
          @click="handleNextClick"
        >
          {{ t('quiz_next') }}
        </button>

        <button
          v-else
          class="btn-repeat"
          @click="resetSelection"
        >
          {{ t('quiz_repeat') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from '@/i18n/lang'
import { quizSession } from '@/stores/quizSession'
import {
  getCurrentQuestion,
  loadQuiz,
  goToNextQuestion,
  isLastQuestion,
  markQuizFinished,
} from '@/services/quizService'
import { sendCompletionToBackend } from '@/services/completionService'

const router = useRouter()
const { t } = useI18n()

const howMany = 10
const selectedIndex = ref<number | null>(null)
const state = ref<'idle' | 'correct' | 'wrong'>('idle')

const prefixLetters = ['A', 'B', 'C', 'D']

const currentQuestion = computed(() => getCurrentQuestion())
const totalQuestions = computed(() => quizSession.questions.length)
const currentIndex = computed(() => quizSession.currentIndex)

const progressPercent = computed(() => {
  if (totalQuestions.value === 0) return 0
  return ((currentIndex.value + 1) / totalQuestions.value) * 100
})

function isSelected(index: number) {
  return selectedIndex.value === index
}

function answerButtonClass(index: number) {
  if (!isSelected(index)) return ''
  return state.value === 'correct' ? 'answer-correct' : 'answer-wrong'
}

function onAnswerClick(index: number) {
  if (state.value !== 'idle') return  // ⛔ BLOKADA ODPOWIEDZI

  if (!currentQuestion.value) return

  selectedIndex.value = index
  state.value =
    index === currentQuestion.value.correctIndex ? 'correct' : 'wrong'
}


function resetSelection() {
  selectedIndex.value = null
  state.value = 'idle'
}

async function handleNextClick() {
  if (isLastQuestion()) {
    markQuizFinished()
    await sendCompletionToBackend()
    router.push({ name: 'bye' })
    return
  }

  goToNextQuestion()
  resetSelection()
}

onMounted(async () => {
  if (!quizSession.questions.length) {
    await loadQuiz(howMany)
  }
})
</script>

<style scoped>
.question-view {
  min-height: 100vh;
  width: 100%;
  display: grid;
  grid-template-columns: 1.7fr 2fr;
}

/* LEWA KOLUMNA */

.left-panel {
  background: linear-gradient(
    135deg,
    var(--color-pink-light),
    var(--color-orange)
  );
  display: flex;
  align-items: stretch;
}

.left-inner {
  padding: 40px 64px 32px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 100%;
}

.counter {
  font-size: 20px;
  font-weight: 600;
  color: #ffffff;
  letter-spacing: 1px;
}

.question-text {
  margin-top: 80px;
  font-size: 36px;
  font-weight: 600;
  color: #ffffff;
}

/* PROGRESS BAR NA DOLE */

.progress-bar {
  width: 80%;
  height: 32px;
  margin: 0 auto;
  border-radius: 6px;
  border: 2px solid #000000;
  position: relative;
  overflow: hidden;
  background: transparent;
}

.progress-bar-fill {
  height: 100%;
  background: linear-gradient(
    135deg,
    var(--color-pink-light),
    var(--color-orange)
  );
  transition: width 0.3s ease;
}

.progress-label {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  font-weight: 600;
}

/* PRAWA KOLUMNA */

.right-panel {
  background: var(--color-dark-bg);
  padding: 40px 64px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

/* LISTA ODPOWIEDZI */

.answers {
  display: flex;
  flex-direction: column;
  gap: 24px;
  margin-top: 24px;
}

.answer-button {
  position: relative;
  width: 100%;
  padding: 18px 32px;
  border-radius: 32px;
  border: 3px solid #ffffff;

  background: var(--color-pink-light);
  color: #ffffff;

  display: flex;
  align-items: center;
  gap: 24px;
  text-align: left;

  font-size: 22px;
}

/* BĄBEL Z LITERĄ A/B/C/D */

.answer-prefix {
  width: 60px;
  height: 60px;
  border-radius: 999px;
  border: 2px solid #ffffff;
  background: radial-gradient(circle at 30% 30%, #ffffff55, var(--color-pink-light));
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 26px;
  font-weight: 700;
  color: #ffffff;
  flex-shrink: 0;
}

.status-img {
  width: 80px;       /* większe niż bąbel — jak w mocku */
  height: 80px;
  object-fit: contain;
  flex-shrink: 0;

  margin-left: -6px; /* wyrównanie względem bąbla */
}


.answer-text {
  flex: 1;
}

/* STANY PRZYCISKU ODPOWIEDZI */

.answer-correct {
  background: linear-gradient(
    135deg,
    var(--color-green-dark),
    var(--color-green-light)
  );
}

.answer-wrong {
  background: linear-gradient(
    135deg,
    var(--color-red-dark),
    var(--color-red-light)
  );
}

/* PRZYCISKI NEXT / REPEAT */

.controls {
  margin-top: 32px;
  display: flex;
  justify-content: flex-end;
  gap: 16px;
}

.btn-next,
.btn-repeat {
  min-width: 180px;
  height: 64px;
  border-radius: 22px;
  border: 3px solid #ffffff;

  font-size: 24px;
  font-weight: 300;
  color: #ffffff;

  display: flex;
  align-items: center;
  justify-content: center;

  cursor: pointer;
}

.btn-next {
  background: linear-gradient(
    135deg,
    var(--color-green-dark),
    var(--color-green-light)
  );
}

.btn-repeat {
  background: linear-gradient(
    135deg,
    var(--color-blue-dark),
    var(--color-blue-light)
  );
}
</style>