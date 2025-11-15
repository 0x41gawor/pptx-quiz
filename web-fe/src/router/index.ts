import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

import W1HelloView from '@/views/W1HelloView.vue'
import W2UserDataView from '@/views/W2UserDataView.vue'
import W3QuestionView from '@/views/W3QuestionView.vue'
import W4ByeView from '@/views/W4ByeView.vue'
import WXFileUploadView from '@/views/WXFileUploadView.vue'
import WXCompletionsView from '@/views/WXCompletionsView.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'hello',
    component: W1HelloView,
  },
  {
    path: '/user-data',
    name: 'user-data',
    component: W2UserDataView,
  },
  {
    path: '/quiz',
    name: 'quiz',
    component: W3QuestionView,
  },
  {
    path: '/bye',
    name: 'bye',
    component: W4ByeView,
  },
  // Admin area
  {
    path: '/admin/file-upload',
    name: 'admin-file-upload',
    component: WXFileUploadView,
  },
  {
    path: '/admin/completions',
    name: 'admin-completions',
    component: WXCompletionsView,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
