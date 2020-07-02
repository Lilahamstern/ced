import Vue from 'vue'
import Router from 'vue-router'
import HomeView from "@/views/HomeView";

Vue.use(Router)

let routes = [
    {
        path: '/',
        name: "Home",
        component: HomeView
    }
]

export default new Router({
    mode: 'history',
    routes
})
