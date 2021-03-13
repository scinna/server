import {HttpError, ScinnaError} from "../types/Error";

type Method = 'GET' | 'get' | 'POST' | 'post' | 'PUT' | 'put' | 'DELETE' | 'delete';

type ApiParameter = {
    url: string;
    data?: object;
    method?: Method;
}

export async function apiCall<T>(token: string|null, params: ApiParameter): Promise<T|HttpError|ScinnaError> {
    let headers: HeadersInit = {};
    let body: BodyInit | null = null;

    if (token) {
        headers.Authorization = "Bearer " + token;
    }

    if (params.data) {
        body = JSON.stringify(params.data);
    }

    const resp = await fetch(params.url, {
        method: params.method ?? 'GET',
        headers,
        body
    })

    if (!resp.ok) {
        try {
            // Scinna error
            return await resp.json()
        } catch {
            // HTTP Error
            return { status: resp.status };
        }
    }

    return await resp.json();
}