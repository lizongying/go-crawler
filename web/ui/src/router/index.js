import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('@/views/HomeView.vue')
        },
        {
            path: '/crawlers',
            name: 'crawlers',
            component: () => import('@/views/CrawlersView.vue')
        },
        {
            path: '/spiders',
            name: 'spiders',
            component: () => import('@/views/SpidersView.vue')
        },
        {
            path: '/jobs',
            name: 'jobs',
            component: () => import('@/views/JobsView.vue')
        },
        {
            path: '/tasks',
            name: 'tasks',
            component: () => import('@/views/TasksView.vue')
        },
        {
            path: '/requests',
            name: 'requests',
            component: () => import('@/views/RequestsView.vue')
        },
        {
            path: '/items',
            name: 'items',
            component: () => import('@/views/ItemsView.vue')
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/LoginView.vue')
        },
        {
            path: '/:pathMatch(.*)*',
            name: '404',
            component: () => import('@/views/404.vue')
        }
    ]
})

export default router
