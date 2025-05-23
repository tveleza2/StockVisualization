import type { StockPrice } from "@/ports/StockPrice";

export function GetStockPrices():Promise<StockPrice[]>{ 
  return new Promise((resolve,reject)=>{
    resolve([
      { time: new Date("2024-01-01"), "price": 150 },
      { time: new Date("2024-01-02"), "price": 152 },
      { time: new Date("2024-01-03"), "price": 148 },
      { time: new Date("2024-01-04"), "price": 155 },
      { time: new Date("2024-01-05"), "price": 157 },
      { time: new Date("2024-01-06"), "price": 153 },
      { time: new Date("2024-01-07"), "price": 158 },
      { time: new Date("2024-01-08"), "price": 160 },
      { time: new Date("2024-01-09"), "price": 162 },
      { time: new Date("2024-01-10"), "price": 159 },
      { time: new Date("2024-01-11"), "price": 161 },
      { time: new Date("2024-01-12"), "price": 163 },
      { time: new Date("2024-01-13"), "price": 165 },
      { time: new Date("2024-01-14"), "price": 167 },
      { time: new Date("2024-01-15"), "price": 170 },
      { time: new Date("2024-01-16"), "price": 168 },
      { time: new Date("2024-01-17"), "price": 171 },
      { time: new Date("2024-01-18"), "price": 173 },
      { time: new Date("2024-01-19"), "price": 175 },
      { time: new Date("2024-01-20"), "price": 172 }
      ]);
})}
