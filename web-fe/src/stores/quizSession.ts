import { reactive } from 'vue'
import type { Language } from '@/models/common'
import type { Question } from '@/models/quiz'

export interface QuizSessionState {
  language: Language | null
  firstname: string
  lastname: string
  phoneNumber: string
  questions: Question[]
  currentIndex: number
  startedAt: string | null
  finishedAt: string | null
}

export const quizSession = reactive<QuizSessionState>({
  language: null,
  firstname: '',
  lastname: '',
  phoneNumber: '',
  questions: [],
  currentIndex: 0,
  startedAt: null,
  finishedAt: null,
})
