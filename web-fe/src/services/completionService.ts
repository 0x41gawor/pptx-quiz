import { quizSession } from '@/stores/quizSession'
import { createCompletion } from '@/api/completionsApi'

export async function sendCompletionToBackend(): Promise<void> {
  if (!quizSession.firstname || !quizSession.lastname || !quizSession.phoneNumber) {
    throw new Error('User data not set')
  }

  const date = quizSession.finishedAt ?? new Date().toISOString()

  await createCompletion({
    firstname: quizSession.firstname,
    lastname: quizSession.lastname,
    phone_number: quizSession.phoneNumber,
    date,
  })
}
