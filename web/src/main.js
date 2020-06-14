// Vuejs
import Vue from 'vue';
import App from './App.vue';
import router from './router';
import devtools from '@vue/devtools';

// GraphQL
import { ApolloClient} from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';
import  {setContext} from "apollo-link-context";
import VueApollo from 'vue-apollo;'

// Custom
import '@/assets/css/tailwind.css';
import i18n from './locales/i18n';
import { library, dom } from '@fortawesome/fontawesome-svg-core';
import { fas } from '@fortawesome/free-solid-svg-icons';


// Devtools
if (process.env.NODE_ENV === 'development') {
  devtools.connect();
}

Vue.config.productionTip = true;
Vue.config.performance = true;

dom.watch();

// Library
library.add(fas);

// GraphQL
const httpLink = new HttpLink({
  uri: "http://localhost:3000/graphql"
});

const middlewareLink = setContext(() => ({
  headers: {
  }
}));

const link = middlewareLink.concat(httpLink);

// Apollo client
const apolloClient = new ApolloClient({
  link,
  cache: new InMemoryCache(),
  connectToDevTools: true
})

const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
});

// Vue
new Vue({
  provide: apolloProvider.provide(),
  router,
  i18n,
  render: h => h(App)
}).$mount('#app');
