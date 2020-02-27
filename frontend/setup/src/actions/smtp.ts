export const ACTION_UPDATE_SMTP = "UPDATE_SMTP";

interface updateSmtpPayload {
    Enabled?:      boolean,
    Hostname?:     string,
    Port?:         string,
    Username?:     string,
    Password?:     string,
    Sender?:       string,
    TestReceiver?: string,
}

export const actionUpdateSmtp = (payload: updateSmtpPayload) => {
    return { type: ACTION_UPDATE_SMTP, payload };
}