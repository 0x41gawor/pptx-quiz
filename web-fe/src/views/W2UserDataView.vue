<template>
  <div class="user-data-view">
    <div class="form-card">

      <h2 class="form-title">
              {{ t('user_data_title').toUpperCase() }}
      </h2>

      <form @submit.prevent="handleNext" class="form-grid">
        <div class="row">
          <label class="label">{{ t('user_data_firstname') }}</label>
          <input v-model="firstname" type="text" required />
        </div>

        <div class="row">
          <label class="label">{{ t('user_data_lastname') }}</label>
          <input v-model="lastname" type="text" required />
        </div>

        <div class="row">
          <label class="label">{{ t('user_data_phone_number') }}</label>
          <input v-model="phoneNumber" type="tel" required />
        </div>
      </form>
    </div>

    <button class="next-btn" @click="handleNext" :disabled="!isValid">
      {{ t('user_data_next').toUpperCase() }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from '@/i18n/lang'
import { initializeSession } from '@/services/quizService'
import { quizSession } from '@/stores/quizSession'

const { t } = useI18n()
const router = useRouter()
const firstname = ref(quizSession.firstname)
const lastname = ref(quizSession.lastname)
const phoneNumber = ref(quizSession.phoneNumber)

onMounted( () => {
  firstname.value = ""
  lastname.value = "" 
  phoneNumber.value = ""
})

const isValid = computed(() =>
  firstname.value.trim().length > 0 &&
  lastname.value.trim().length > 0 &&
  phoneNumber.value.trim().length > 0
)

function handleNext() {
  if (!quizSession.language) {
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
  width: 100%;
  background: url('/backgrounds/2.png');
  background-size: cover;
  background-position: right center;
  background-repeat: no-repeat;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding-left: 6vw;
  gap: 6vw;
}

/* ------- BIG GRADIENT CARD ------- */

.form-card {
  width: 900px;
  padding: 60px 80px;
  border-radius: 80px;
  border: 6px solid white;

  background: linear-gradient(
    135deg,
    color-mix(in srgb, var(--color-pink-light) 80%, transparent),
    color-mix(in srgb, var(--color-orange) 80%, transparent)
  );

  display: flex;
  flex-direction: column;
  align-items: center;
}

.form-title {
  color: white;
  font-size: 56px;
  font-weight: 100;
  text-align: center;
  line-height: 1.2;
}

/* ------- FORM GRID ------- */

.form-grid {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 40px;
}

.row {
  display: grid;
  grid-template-columns: 1fr 1.6fr;
  align-items: center;

  padding: 16px 28px;
  border-radius: 20px;
  border: 3px solid white;

  background: transparent;
}

.label {
  color: white;
  font-size: 24px;
  font-weight: 300;
}

input {
  width: 100%;
  height: 52px;
  border-radius: 14px;
  border: none;

  padding: 0 16px;
  font-size: 22px;

  background: white;
  color: black;
}

/* ------- NEXT BUTTON (outside card) ------- */

.next-btn {
  width: 323px;
  height: 100px;

  font-size: 28px;
  font-weight: 300;
  margin-top: 62vh;
  margin-right: 5vh;

  border-radius: 20px;
  border: 3px solid white;

  background: linear-gradient(
    135deg,
    var(--color-green-dark),
    var(--color-green-light)
  );

  color: white;
  cursor: pointer;

  display: flex;
  align-items: center;
  justify-content: center;

  transition: transform 0.2s ease;
}

.next-btn:hover {
  transform: scale(1.03);
}

.next-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
  transform: none;
}

.next-btn:disabled:hover {
  transform: none;
}

</style>
