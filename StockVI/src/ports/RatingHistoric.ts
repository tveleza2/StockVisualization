export interface RatingsHistoric{ 
    stock: string,
    company: string, 
    broker: string, 
    action: string, 
    rating: string,
    target?: number,
    recommended?: boolean
}

