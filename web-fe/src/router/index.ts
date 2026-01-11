import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

import W1HelloView from '@/views/W1HelloView.vue'
import W2UserDataView from '@/views/W2UserDataView.vue'
import W3QuestionView from '@/views/W3QuestionView.vue'
import W4ByeView from '@/views/W4ByeView.vue'
import WXFileUploadView from '@/views/WXFileUploadView.vue'
import WXCompletionsView from '@/views/WXCompletionsView.vue'
import WYIVideoView from '@/views/WYVideoView.vue'
import WYVideoEndedView from '@/views/WYVideoEndedView.vue'


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
    path: '/admin',
    name: 'admin',
    component: WXCompletionsView,
  },
  {
    path: '/plik',
    name: 'admin-file',
    component: WXFileUploadView,
  },
  {
    path: '/wyniki',
    name: 'admin-completions',
    component: WXCompletionsView,
  },
  // Video area
  {
  path: '/video',
  name: 'video',
  component: WYIVideoView,
  },
  {
  path: '/video-ended',
  name: 'video-ended',
  component: WYVideoEndedView,
},

]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
