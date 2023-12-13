import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
        },
        {
            path: '/image/:id',
            name: 'image',
            component: () => import('../views/ImageView.vue'),
            // beforeEnter(to) {
            //     // console.log('Accessing route --> ', to.params)
            //     // eslint-disable-next-line no-constant-condition
            //     if (!to.params.id) {
            //       return { name: 'home' }
            //     }
            //   },
        },
        {
            path: '/about',
            name: 'about',
            component: () => import('../views/AboutView.vue'),
        },
        // {
        //   path: '/:pathMatch(.*)*',
        //   name: '404',
        //   component: () => import('../views/NotFoundView.vue')
        // },
    ]
})

export default router
