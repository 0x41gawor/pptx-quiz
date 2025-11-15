import type { ApiQuizResponse } from '@/models/quiz'
import type { Language } from '@/models/common'
import { httpClient } from './httpClient'

export async function getQuiz(lang: Language, howMany: number): Promise<ApiQuizResponse> {
  const params = new URLSearchParams({
    lang,
    howMany: String(howMany),
  }).toString()

  return httpClient.get<ApiQuizResponse>(`/quiz?${params}`)
}

// File endpoints dla quiz CSV – użyjemy później w WX-file-upload
export async function downloadQuizFile(): Promise<Blob> {
  const url = 'http://192.46.236.119:3456/api/v1/quiz/file'
  const response = await fetch(url, { method: 'GET' })
  if (!response.ok) {
    throw new Error('Failed to download quiz file')
  }
  return await response.blob()
}

export async function uploadQuizFile(file: File): Promise<void> {
  const url = 'http://192.46.236.119:3456/api/v1/quiz/file'
  const formData = new FormData()
  formData.append('file', file)

  const response = await fetch(url, {
    method: 'POST',
    body: formData,
  })

  if (!response.ok) {
    const text = await response.text().catch(() => '')
    throw new Error(`Failed to upload quiz file: ${text || response.statusText}`)
  }
}
