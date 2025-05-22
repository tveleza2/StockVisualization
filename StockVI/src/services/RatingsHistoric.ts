import type { RatingsHistoric } from "@/ports/RatingHistoric"


export async function GetRatingsHistoric():Promise<RatingsHistoric[]>{ 
  const res = await fetch('http://localhost:8080/rating-historics/')
  if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
  const apiData = await res.json();
  return apiData.map((item: any): RatingsHistoric => ({
      stock: item.ticker,
      company: item.company, // or you could hardcode a map if needed
      broker: item.brokerage,
      action: item.action,
      rating: item.rating_to,
      target: parseFloat(item.target_to.replace('$', '')) || undefined
    }))}



