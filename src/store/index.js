import { createStore } from "vuex"


const store = createStore({
    state:{
        user:{
        }
    },
    mutations: {
        setUserInfo(state, user) {
            state.user = user
        }
    }
})


export default store