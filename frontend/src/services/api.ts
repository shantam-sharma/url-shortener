import type {
    CreateURLRequest,
    CreateURLResponse,
} from "../types/api";

const API_BASE_URL =
    import.meta.env.VITE_API_URL || "http://127.0.0.1:8080";

export async function createShortURL(
    request: CreateURLRequest
): Promise<CreateURLResponse> {

    const response = await fetch(`${API_BASE_URL}/api/v1/urls`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(request),
    });

    if (!response.ok) {
        const error = await response.text();
        throw new Error(error);
    }

    return response.json();
}
