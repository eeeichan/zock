import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount('#app');

// new Vue({
//   el: '#app',
//   data: {
//     users: [
//       {
//         id: 1,
//         name: '鈴木太郎',
//         email: 'suzukitaro@example.com'
//       },
//       {
//         id: 2,
//         name: '佐藤二郎',
//         email: 'satoujiro@example.com'
//       },
//       {
//         id: 3,
//         name: '田中三郎',
//         email: 'tanakasaburo@example.com'
//       },
//       {
//         id: 4,
//         name: '山本四郎',
//         email: 'yamamotoshiro@example.com'
//       },
//       {
//         id: 5,
//         name: '高橋五郎',
//         email: 'takahashigoro@example.com'
//       }
//     ]
//   }
// });
