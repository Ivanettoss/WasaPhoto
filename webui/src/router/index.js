import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
    history: createWebHashHistory(
        import.meta.env.BASE_URL),
        routes: [

        {
            path: '/',
            component: LoginView
        },
        
       {
            path: '/profile/:username',
            component: ProfileView
        },
        
        {
            path: '/home',
            component: HomeView
        },
         
    ]})

        export default router