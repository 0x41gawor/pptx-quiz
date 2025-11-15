export interface CompletionPostPayload {
  firstname: string
  lastname: string
  phone_number: string
  date: string // ISO string
}

export interface Completion {
  id: number
  firstname: string
  lastname: string
  phone_number: string
  date: string // ISO string from backend
}

export interface CompletionQueryParams {
  query?: string
}
