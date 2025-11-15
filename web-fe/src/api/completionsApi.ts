import type {
  Completion,
  CompletionPostPayload,
} from '@/models/completion'
import { httpClient } from './httpClient'

export async function createCompletion(
  payload: CompletionPostPayload,
): Promise<Completion> {
  return httpClient.post<Completion>('/completions', payload)
}

export interface CompletionsListResponse {
  completions: Completion[]
  // jeżeli backend zwraca inaczej – dopasujemy później
}

export async function listCompletions(query?: string): Promise<Completion[]> {
  const q = query ? `?query=${encodeURIComponent(query)}` : ''
  // zakładam, że backend zwraca prostą listę – w razie czego poprawimy
  return httpClient.get<Completion[]>(`/completions${q}`)
}
