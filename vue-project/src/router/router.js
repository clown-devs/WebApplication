import Main from '@/pages/Main'
import Login from '@/pages/Login'
import Calendar from '@/pages/Calendar'
import Clients from '@/pages/Clients'
import MeetingResults from '@/pages/MeetingResults'
import {createRouter, createWebHistory} from 'vue-router'
import api from "@/api";

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

router.beforeEach(async (to, from, next) => {
    document.title = to.name
    if (to.name !== 'Вход' && !await api.isValidToken()) next('/login')
    else next()
})

export default router;