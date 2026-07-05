export interface CreateURLRequest {
    url: string;
    alias?: string;
}

export interface CreateURLResponse {
    original_url: string;
    short_code: string;
    short_url: string;
}
