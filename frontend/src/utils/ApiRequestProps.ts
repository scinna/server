import {ScinnaError} from "../types/Error";

export type HttpMethod = 'GET' | 'get' | 'POST' | 'post' | 'PUT' | 'put' | 'DELETE' | 'delete';

export type HttpStatus<T> =
    { status: 'pending' }
    | { status: 'success', data: T }
    | { status: 'error', error: ScinnaError }