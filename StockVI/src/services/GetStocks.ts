import type { Stock } from "@/ports/Stock";



export async function GetStocks():Promise<Stock[]>{
    const res = await fetch('http://localhost:8080/stocks/')
    if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
    const apiData = await res.json();
    return apiData
}
