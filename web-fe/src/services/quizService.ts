import { quizSession } from '@/stores/quizSession'
import type { Language } from '@/models/common'
import type { Question } from '@/models/quiz'
import { getQuiz } from '@/api/quizApi'

// inicjalizacja sesji po wyborze języka i danych użytkownika
export function initializeSession(params: {
  language: Language
  firstname: string
  lastname: string
  phoneNumber: string
}) {
  quizSession.language = params.language
  quizSession.firstname = params.firstname
  quizSession.lastname = params.lastname
  quizSession.phoneNumber = params.phoneNumber
  quizSession.currentIndex = 0
  quizSession.questions = []
  quizSession.startedAt = new Date().toISOString()
  quizSession.finishedAt = null
}

// pobranie pytań z backendu
export async function loadQuiz(howMany: number): Promise<void> {
  if (!quizSession.language) {
    throw new Error('Language not set in quiz session')
  }

  const apiResponse = await getQuiz(quizSession.language, howMany)

  const questions: Question[] = apiResponse.Questions.map((q) => ({
    content: q.Content,
    answers: q.Answers.map((a) => ({ content: a.Content })),
    correctIndex: q.Correct,
  }))

  quizSession.questions = questions
  quizSession.currentIndex = 0
}

// helpery dla widoków
export function getCurrentQuestion(): Question | null {
  const { questions, currentIndex } = quizSession
  if (!questions.length) return null
  return questions[currentIndex] ?? null
}

export function isLastQuestion(): boolean {
  return quizSession.currentIndex >= quizSession.questions.length - 1
}

export function goToNextQuestion(): void {
  if (!isLastQuestion()) {
    quizSession.currentIndex += 1
  }
}

export function markQuizFinished(): void {
  quizSession.finishedAt = new Date().toISOString()
}
