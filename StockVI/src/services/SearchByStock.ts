import type { RatingsHistoric } from '@/ports/RatingHistoric'

export async function searchRatingsHistoric(query: string): Promise<RatingsHistoric[]> {
  const response = await fetch(`http://localhost:8080/rating-historics/${encodeURIComponent(query)}`)
  if (!response.ok) throw new Error('Failed to fetch search results')
  return await response.json()
}