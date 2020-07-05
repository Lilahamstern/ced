import Vue from 'vue'
import VueRouter from 'vue-router'
import Projects from '../views/Projects.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Projects',
    component: Projects,
    meta: {
      breadCrumb: "Projects"
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
