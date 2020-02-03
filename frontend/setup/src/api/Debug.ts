const IN_DEV = true;

export function GetUrl(endpoint: string): string {
    return (IN_DEV ? "https://example.com" : "") + endpoint;
}