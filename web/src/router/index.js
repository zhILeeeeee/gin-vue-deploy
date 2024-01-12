
import { createRouter, createWebHashHistory } from 'vue-router'


export const constantRouters = [
    {
        path: '/docker',
        component: () => import('../views/docker/index.vue'),
        hidden: true,
    },
    {
        path: '/harbor',
        component: () => import('../views/harbor/index.vue')
    },
    {
        path: '/',
        component: () => import('../layout/index.vue')
    }
]

const createCustomerRouter = () =>
    createRouter({
        history: createWebHashHistory(),
        scrollBehavior: () => ({ y: 0 }),
        routes: constantRouters
    })

const router = createCustomerRouter()

export function resetRouter() {
    const newRouter = createCustomerRouter()
    router.matcher = newRouter.matcher
}

export default router