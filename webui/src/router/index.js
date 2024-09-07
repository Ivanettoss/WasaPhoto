import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

import LoginView from '../Views/LoginView.vue'

const router = createRouter({
    history: createWebHashHistory(
        import.meta.env.BASE_URL),
        routes: [{
            path: '/',
            redirect: '/login'
        },

        {
            path: '/login',
            component: LoginView
        },
        
        {
            path: '/profile/:username',
            component: ProfileView
        },
        
        {
            path: '/home',
            component: HomeView
        }
         
    ]})

        export default router