import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import '@/assets/css/tailwind.css';
import devtools from '@vue/devtools';
import i18n from './locales/i18n';
import { library, dom } from '@fortawesome/fontawesome-svg-core';
import { fas } from '@fortawesome/free-solid-svg-icons';

if (process.env.NODE_ENV === 'development') {
  devtools.connect();
}

library.add(fas);
dom.watch();

Vue.config.productionTip = true;
Vue.config.performance = true;

new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app');
