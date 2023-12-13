import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VWave from 'v-wave'
import {
    createSharedElementDirective,
    SharedElementRouteGuard,
    SharedElementDirective
} from 'v-shared-element'
import VueDragscroll from "vue-dragscroll";


const app = createApp(App)
app.use(SharedElementDirective)
app.use(router)
router.beforeEach(SharedElementRouteGuard)
app.use(VWave,
    {
        color: 'white',
        initialOpacity: 0.5,
        easing: 'ease-in',
        duration: 0.3
    })
app.use(VueDragscroll)

app.mount('#app')
