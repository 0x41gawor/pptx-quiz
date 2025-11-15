import { ref, computed } from 'vue'
import type { Language } from '@/models/common'

const currentLanguage = ref<Language>('pl')

const messages = {
  pl: {
    hello_title: 'Test wiedzy',
    hello_subtitle: 'Wybierz wersję językową',
    user_data_title: 'Podaj swoje dane',
    user_data_next: 'Dalej',
    quiz_next: 'Dalej',
    quiz_repeat: 'Powtórz',
    bye_title: 'Dziękujemy',
    bye_subtitle: 'Proszę przejść do recepcji po dalsze instrukcje',
  },
  en: {
    hello_title: 'Knowledge test',
    hello_subtitle: 'Choose your language version',
    user_data_title: 'Enter your data',
    user_data_next: 'Next',
    quiz_next: 'Next',
    quiz_repeat: 'Repeat',
    bye_title: 'Thank you',
    bye_subtitle: 'Please go to the reception for further instructions',
  },
}

type MessageKey = keyof typeof messages.pl

export function useI18n() {
  const lang = computed(() => currentLanguage.value)

  function setLanguage(newLang: Language) {
    currentLanguage.value = newLang
  }

  function t(key: MessageKey) {
    return messages[currentLanguage.value][key]
  }

  return {
    lang,
    setLanguage,
    t,
  }
}
