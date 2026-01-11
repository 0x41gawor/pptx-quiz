import { ref, computed } from 'vue'
import type { Language } from '@/models/common'

const currentLanguage = ref<Language>('pl')

const messages = {
  pl: {
    hello_title: 'Test wiedzy',
    hello_subtitle: 'Wybierz wersję językową',
    user_data_title: 'Podaj swoje dane',
    user_data_firstname: 'Imię',
    user_data_lastname: 'Nazwisko',
    user_data_phone_number: 'Numer telefonu',
    user_data_next: 'Dalej',
    quiz_next: 'Dalej',
    quiz_repeat: 'Powtórz',
    question: 'Pytanie',
    bye_title: 'Dziękujemy',
    bye_subtitle: 'Proszę przejść do recepcji po dalsze instrukcje',
    video_title: 'Szkolenie Wstępne',
    video_ended_title: 'Dziękujemy',
    video_ended_subtitle: 'Prosimy o wypełnienie testu na tablecie.',
  },
  en: {
    hello_title: 'Knowledge test',
    hello_subtitle: 'Choose your language version',
    user_data_title: 'Fill in the fields below',
    user_data_next: 'Next',
    user_data_firstname: 'Firstname',
    user_data_lastname: 'Lastname',
    user_data_phone_number: 'Phone number',
    quiz_next: 'Next',
    quiz_repeat: 'Repeat',
    question: 'Question',
    bye_title: 'Thank you',
    bye_subtitle: 'Please go to the reception for further instructions',
    video_title: 'Introductory Training',
    video_ended_title: 'Thank you',
    video_ended_subtitle: 'Please complete the test on the tablet.',
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
