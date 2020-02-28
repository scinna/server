export const ACTION_UPDATE_USER = "UPDATE_USER";

interface updateUserPayload {
    Username?: string,
    Email?: string,
    Password?: string,
}

export const actionUpdateUser = (payload: updateUserPayload) => {
    return { type: ACTION_UPDATE_USER, payload };
}