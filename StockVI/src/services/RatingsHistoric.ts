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




export function GetSymbolRatingsHistoric(symbol:string):Promise<RatingsHistoric[]>{ 
  return new Promise((resolve,reject)=>{
    resolve([
  { stock: symbol, company: 'Apple Inc.', broker: 'JP Morgan', action: 'Reiterated', rating: 'Buy', target: 190.5 },
  { stock: symbol, company: 'Apple Inc.', broker: 'Goldman Sachs', action: 'Target Raised', rating: 'Buy', target: 195.0 },
  { stock: symbol, company: 'Apple Inc.', broker: 'Morgan Stanley', action: 'Reiterated', rating: 'Hold', target: 185.25 },
  { stock: symbol, company: 'Apple Inc.', broker: 'Barclays', action: 'Downgraded', rating: 'Sell', target: 175.0 },
  { stock: symbol, company: 'Apple Inc.', broker: 'UBS', action: 'Initiated', rating: 'Buy', target: 200.75 },
  { stock: symbol, company: 'Apple Inc.', broker: 'BofA Securities', action: 'Reiterated', rating: 'Buy', target: 192.3 },
  { stock: symbol, company: 'Apple Inc.', broker: 'Citigroup', action: 'Target Lowered', rating: 'Neutral', target: 178.9 },
  { stock: symbol, company: 'Apple Inc.', broker: 'Wells Fargo', action: 'Upgraded', rating: 'Buy', target: 198.6 },
  { stock: symbol, company: 'Apple Inc.', broker: 'HC Wainwright', action: 'Initiated', rating: 'Hold', target: 182.0 },
  { stock: symbol, company: 'Apple Inc.', broker: 'Raymond James', action: 'Target Raised', rating: 'Buy', target: 205.4 }
]
);
})}

