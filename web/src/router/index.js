import Vue from 'vue'
import VueRouter from 'vue-router'
import Projects from '../views/Projects.vue'
import Project from "@/views/Project";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Projects,
        meta: {
            breadCrumb: "Home"
        },
    },
    {
        path: '/project/:id',
        name: 'project',
        component: Project,
        meta: {
            breadCrumb: "Project"
        }
    },
    {
        path: '/about',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
        meta: {
            breadCrumb: "About"
        },
        children: [
            {
                path: 'author',
                component: Projects,
                meta: {
                    breadCrumb: 'Author'
                }
            }
        ]
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
