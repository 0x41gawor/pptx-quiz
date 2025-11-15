<template>
  <div class="question-view">
    <div class="left-panel">
      <div v-if="currentQuestion">
        <p class="counter">
          {{ currentIndex + 1 }} / {{ totalQuestions }}
        </p>
        <h2 class="question-text">
          {{ currentQuestion.content }}
        </h2>

        <div class="progress-bar">
          <div class="progress-bar-fill" :style="{ width: progressPercent + '%' }"></div>
          <span class="progress-label">
            {{ Math.round(progressPercent) }}%
          </span>
        </div>
      </div>
    </div>

    <div class="right-panel" v-if="currentQuestion">
      <div class="answers">
        <button
          v-for="(answer, index) in currentQuestion.answers"
          :key="index"
          class="answer-button"
          :class="answerButtonClass(index)"
          @click="onAnswerClick(index)"
        >
          <span class="answer-prefix">{{ prefixLetters[index] }}</span>
          <span class="answer-text">{{ answer.content }}</span>
        </button>
      </div>

      <div class="controls" v-if="state !== 'idle'">
        <button
          v-if="state === 'correct'"
          class="primary"
          @click="handleNextClick"
        >
          {{ t('quiz_next') }}
        </button>
        <button
          v-else
          class="secondary"
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

const howMany = 10 // na razie stała; możemy później parametrYZować
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

function answerButtonClass(index: number) {
  if (selectedIndex.value === null) return ''
  if (index !== selectedIndex.value) return ''
  return state.value === 'correct' ? 'answer-correct' : 'answer-wrong'
}

function onAnswerClick(index: number) {
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
  display: grid;
  grid-template-columns: 2fr 3fr;
}

.left-panel {
  padding: 2rem;
  background: linear-gradient(to bottom, #0f172a, #1f2937);
}

.right-panel {
  padding: 2rem;
  background: #020617;
}

.counter {
  font-size: 1.2rem;
  color: var(--color-text-muted);
}

.question-text {
  margin-top: 0.75rem;
}

.progress-bar {
  margin-top: 1.5rem;
  position: relative;
  height: 1rem;
  border-radius: 999px;
  background: rgba(15, 23, 42, 0.8);
  overflow: hidden;
}

.progress-bar-fill {
  height: 100%;
  background: var(--color-primary);
  transition: width 0.3s ease;
}

.progress-label {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.75rem;
}

.answers {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.answer-button {
  width: 100%;
  border-radius: 999px;
  border: 1px solid rgba(148, 163, 184, 0.5);
  padding: 0.75rem 1rem;
  background: #020617;
  color:aliceblue;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  text-align: left;
}

.answer-prefix {
  width: 2rem;
  height: 2rem;
  border-radius: 999px;
  border: 1px solid rgba(148, 163, 184, 0.6);
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.answer-correct {
  background: rgba(34, 197, 94, 0.15);
  border-color: var(--color-success);
}

.answer-wrong {
  background: rgba(239, 68, 68, 0.15);
  border-color: var(--color-error);
}

.controls {
  margin-top: 1.5rem;
  display: flex;
  justify-content: flex-end;
  gap: 0.75rem;
}

button.primary {
  border-radius: 999px;
  border: none;
  padding: 0.6rem 1.4rem;
  background: var(--color-primary);
  color: #000;
  font-weight: 600;
}

button.secondary {
  border-radius: 999px;
  border: 1px solid rgba(148, 163, 184, 0.7);
  padding: 0.6rem 1.4rem;
  background: transparent;
  color: var(--color-text-main);
}
</style>
