import type { StockPrice } from "@/ports/StockPrice";

export function GetStockPrices():Promise<StockPrice[]>{ 
  return new Promise((resolve,reject)=>{
    resolve([
        { time: new Date("2024-01-01"), "price": 150 },
        { time: new Date("2024-01-02"), "price": 152 },
        { time: new Date("2024-01-03"), "price": 148 }
      ]);
})}
