import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import '@/assets/css/tailwind.css';
import devtools from '@vue/devtools';

if (process.env.NODE_ENV === 'development') {
  devtools.connect();
}

Vue.config.productionTip = true;
Vue.config.performance = true;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
