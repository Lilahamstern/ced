import 'material-design-icons-iconfont/dist/material-design-icons.css'
import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    dark: true,
      options: {
        customProperties: true,
      },
    themes: {
      dark: {
        primary: '#2196f3',
        secondary: '#009688',
        accent: '#3f51b5',
        error: '#f44336',
        warning: '#ffeb3b',
        info: '#ff9800',
        success: '#4caf50',
      },
    },
  },
  icons: {
    iconfont: 'md',
  },
});
