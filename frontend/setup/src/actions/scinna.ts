export const ACTION_UPDATE_SCINNA = "UPDATE_SCINNA";

interface updateScinnaPayload {
    Registration?: string,
    IPHeader?: string,
    RateLimiting?: string,
    PicturePath?: string,
    WebURL?: string,
}

export const actionUpdateScinna = (payload: updateScinnaPayload) => {
    return { type: ACTION_UPDATE_SCINNA, payload };
}