import {createRouter, createWebHashHistory} from 'vue-router'
import StreamView from '../views/StreamView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView},
		{path: '/myStream', component: StreamView},
		{path: '/profile', component: ProfileView},
	]
})

export default router
