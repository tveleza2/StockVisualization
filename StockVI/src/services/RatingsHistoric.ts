import type { RatingsHistoric } from "@/ports/RatingHistoric"

export function GetRatingsHistoric():RatingsHistoric[]{

    // WRITE FUNCTION TO RETRIEVE THE RATINGS FROM THE BACKEND API
    return [
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'NVDA', company: 'NVIDIA', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
        { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
    ]
}