import i18n from "i18n-js";

export const displayDate = (apiDate: string|null): string|null => {
    const parsedDate = apiDate ? new Date(Date.parse(apiDate)) : null;
    if (!parsedDate) return null;

    return parsedDate.toLocaleTimeString() + " " + i18n.t('my_profile.tokens.on') + " " + parsedDate.toLocaleDateString();
};