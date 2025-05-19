import type { RatingsHistoric } from "@/ports/RatingHistoric"


export function GetRatingsHistoric():Promise<RatingsHistoric[]>{ 
  return new Promise((resolve,reject)=>{
    resolve([
  { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Reiterated', rating: 'Buy' },
  { stock: 'NVDA', company: 'NVIDIA', broker: 'JP Morgan', action: 'Target Lowered', rating: 'Sell' },
  { stock: 'RMTI', company: 'Rockwell Medical', broker: 'HC Wainwright', action: 'Target lowered', rating: 'Neutral' },
  { stock: 'TSLA', company: 'Tesla Inc.', broker: 'Goldman Sachs', action: 'Initiated', rating: 'Buy' },
  { stock: 'AMZN', company: 'Amazon.com Inc.', broker: 'Morgan Stanley', action: 'Reiterated', rating: 'Buy' },
  { stock: 'MSFT', company: 'Microsoft Corp.', broker: 'Barclays', action: 'Downgraded', rating: 'Hold' },
  { stock: 'GOOGL', company: 'Alphabet Inc.', broker: 'UBS', action: 'Upgraded', rating: 'Buy' },
  { stock: 'META', company: 'Meta Platforms', broker: 'BofA Securities', action: 'Reiterated', rating: 'Buy' },
  { stock: 'INTC', company: 'Intel Corp.', broker: 'Citigroup', action: 'Target Raised', rating: 'Neutral' },
  { stock: 'NFLX', company: 'Netflix Inc.', broker: 'Wells Fargo', action: 'Initiated', rating: 'Sell' }
]
);
})}




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

