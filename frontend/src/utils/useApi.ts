import {isScinnaError, ScinnaError} from "../types/Error";
import {DependencyList, useState}   from "react";
import useAsyncEffect               from "use-async-effect";
import {useToken}                   from "./TokenProvider";

type Method = 'GET' | 'get' | 'POST' | 'post' | 'PUT' | 'put' | 'DELETE' | 'delete';

type ApiParameter = {
    url: string;
    data?: object;
    method?: Method;
}

export async function apiCall<T>(token: string | null, params: ApiParameter): Promise<T | ScinnaError> {
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
            return {...await resp.json(), status: resp.status}
        } catch {
            // HTTP Error
            return {Message: 'HTTP Error: ' + resp.status, ErrCode: -1, status: resp.status};
        }
    }

    return await resp.json();
}

export type ApiResponse<T> =
    { status: 'pending' }
    | { status: 'success', data: T }
    | { status: 'error', error: ScinnaError };

export function useApiCall<T>(params: ApiParameter, deps: DependencyList = []) {
    const {loaded, token, logout} = useToken();
    const [apiResponse, setApiResponse] = useState<ApiResponse<T>>({status: 'pending'});

    useAsyncEffect(async () => {
        if (!loaded) {
            return;
        }

        setApiResponse({status: 'pending'});
        const data = await apiCall<T>(token, params);

        if (isScinnaError(data)) {
            if ((data as ScinnaError).status === 401) {
                logout();
            }
            setApiResponse({status: 'error', error: data as ScinnaError})
            return;
        }

        setApiResponse({ status: 'success', data: data as T })
    }, [token, params.url, ...deps]);

    return apiResponse;
}