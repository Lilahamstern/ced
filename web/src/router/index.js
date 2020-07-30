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
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
