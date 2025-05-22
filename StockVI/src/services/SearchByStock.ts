import type { RatingsHistoric } from '@/ports/RatingHistoric'

export async function searchRatingsHistoric(query: string): Promise<RatingsHistoric[]> {
    const url = query.length===0?'http://localhost:8080/rating-historics/':`http://localhost:8080/rating-historics/${encodeURIComponent(query.toUpperCase())}`
    const response = await fetch(url)
    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
    const apiData = await response.json();
    return apiData.map((item: any): RatingsHistoric => ({
        stock: item.ticker,
        company: item.company, // or you could hardcode a map if needed
        broker: item.brokerage,
        action: item.action,
        rating: item.rating_to,
        target: parseFloat(item.target_to.replace('$', '')) || undefined
    }))}