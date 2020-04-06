export const ACTION_UPDATE_SCINNA = "UPDATE_SCINNA";

interface updateScinnaPayload {
    RegistrationAllowed?: string,
    HeaderIPField?: string,
    RateLimiting?: number,
    MediaPath?: string,
    WebURL?: string,
}

export const actionUpdateScinna = (payload: updateScinnaPayload) => {
    return { type: ACTION_UPDATE_SCINNA, payload };
}