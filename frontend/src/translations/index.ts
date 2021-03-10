import fr from './fr';
import en from './en';

import i18n from 'i18n-js';

export const init = () => {
    i18n.translations = { fr, en };
    i18n.locale = 'en';// @TODO: Find local from browser;
    i18n.fallbacks = true;
}
