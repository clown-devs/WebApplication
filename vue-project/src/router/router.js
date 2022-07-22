import Main from '@/pages/Main'
import Login from '@/pages/Login'
import Calendar from '@/pages/Calendar'
import Clients from '@/pages/Clients'
import MeetingResults from '@/pages/MeetingResults'
import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    { 
        path: '/', component: Main
    },

    { 
        path: '/login', component: Login
    },

    { 
        path: '/calendar', component: Calendar
    },

    { 
        path: '/results', component: MeetingResults
    },

    { 
        path: '/clients', component: Clients
    }
]

const router = createRouter({
    routes,
    history: createWebHistory()
})

export default router;