<template>
  <div class="user-data-view">
    <div class="form-card">
      <h2>{{ t('user_data_title') }}</h2>

      <form @submit.prevent="handleNext">
        <label>
          <span>First name / Imię</span>
          <input v-model="firstname" type="text" required />
        </label>

        <label>
          <span>Last name / Nazwisko</span>
          <input v-model="lastname" type="text" required />
        </label>

        <label>
          <span>Phone number / Numer telefonu</span>
          <input v-model="phoneNumber" type="tel" required />
        </label>

        <div class="actions">
          <button type="submit" class="primary">
            {{ t('user_data_next') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from '@/i18n/lang'
import { initializeSession } from '@/services/quizService'
import { quizSession } from '@/stores/quizSession'

const router = useRouter()
const { t } = useI18n()

const firstname = ref(quizSession.firstname)
const lastname = ref(quizSession.lastname)
const phoneNumber = ref(quizSession.phoneNumber)

async function handleNext() {
  if (!quizSession.language) {
    // fallback – wróć do wyboru języka
    router.push({ name: 'hello' })
    return
  }

  initializeSession({
    language: quizSession.language,
    firstname: firstname.value.trim(),
    lastname: lastname.value.trim(),
    phoneNumber: phoneNumber.value.trim(),
  })

  router.push({ name: 'quiz' })
}
</script>

<style scoped>
.user-data-view {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.form-card {
  width: min(480px, 100%);
  background: linear-gradient(
      to bottom right,
      rgba(34, 197, 94, 0.15),
      rgba(15, 23, 42, 0.95)
    );
  padding: 2rem 2.5rem;
  border-radius: 1.5rem;
}

label {
  display: flex;
  flex-direction: column;
  margin-top: 1rem;
  gap: 0.4rem;
}

span {
  font-size: 0.9rem;
  color: var(--color-text-muted);
}

input {
  padding: 0.6rem 0.8rem;
  border-radius: 0.75rem;
  border: 1px solid rgba(148, 163, 184, 0.5);
  background: rgba(15, 23, 42, 0.8);
  color: var(--color-text-main);
}

.actions {
  margin-top: 1.5rem;
  display: flex;
  justify-content: flex-end;
}

button.primary {
  border: none;
  border-radius: 999px;
  padding: 0.6rem 1.4rem;
  background: var(--color-primary);
  color: #000;
  font-weight: 600;
}
</style>
