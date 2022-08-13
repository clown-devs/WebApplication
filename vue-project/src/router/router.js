import Main from '@/pages/Main'
import Login from '@/pages/Login'
import Calendar from '@/pages/Calendar'
import Clients from '@/pages/Clients'
import MeetingResults from '@/pages/MeetingResults'
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    { 
        path: '/', 
        component: Main,
        name: 'Главная'
    },

    { 
        path: '/login', 
        component: Login,
        name: 'Вход'
    },

    { 
        path: '/calendar',
        component: Calendar,
        name: 'Календарь'
    },

    { 
        path: '/results',
        component: MeetingResults,
        name: 'Отчеты'
    },

    { 
        path: '/clients', 
        component: Clients,
        name: 'Клиенты'
    }
]

const router = createRouter({
    routes,
    history: createWebHistory()
})

router.beforeEach((to, from, next) => {
    document.title = to.name
    next()
})

export default router;