import type { Language } from './common'

export interface ApiAnswer {
  Content: string
}

export interface ApiQuestion {
  Content: string
  Answers: ApiAnswer[]
  Correct: number // index w tablicy Answers
}

export interface ApiQuizResponse {
  Questions: ApiQuestion[]
}

// Model używany w UI (bardziej "ts-owy", ale prawie 1:1)
export interface Answer {
  content: string
}

export interface Question {
  content: string
  answers: Answer[]
  correctIndex: number
}

export interface Quiz {
  questions: Question[]
}

export interface QuizRequestParams {
  lang: Language
  howMany: number
}
