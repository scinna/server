export const ACTION_UPDATE_DATABASE = "UPDATE_DATABASE";

interface updateDatabasePayload {
    DBMS?:     string,
    Hostname?: string,
    Port?:     string,
    Username?: string,
    Password?: string,
    Database?: string,
    Path?:     string,
}

export const actionUpdateDatabase = (payload: updateDatabasePayload) => {
    return { type: ACTION_UPDATE_DATABASE, payload };
}