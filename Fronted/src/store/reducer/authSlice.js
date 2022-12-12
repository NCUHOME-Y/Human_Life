import { createSlice } from "@reduxjs/toolkit";

export const authSlice = createSlice({
    name : 'auth',
    initialState : () => {
        
        const token = localStorage.getItem('token')
        if(!token){
            return{
                islogged : false,
                token : null,
                user : null
            }   
        }else{
            return{
                islogged :true,
                token : token,
                user : JSON.stringify(localStorage.getItem('user'))
            }
            
        }
},
    reducers : {
        login(state,action){
            state.islogged = true
            state.token = action.payload.token
            state.user = action.payload.user
            localStorage.setItem('token',state.token)
            localStorage.setItem('user',JSON.stringify(state.user))
            localStorage.setItem('isLogged',state.islogged)
        },
        logout(state,action){   
            state.islogged = false
            state.token = null
            state.user = null
            localStorage.removeItem('token')
            localStorage.removeItem('user') 
            localStorage.removeItem('isLogged')
            localStorage.removeItem('isSet')
            localStorage.removeItem('exsit') 
        }
    }
})

export const {login,logout} = authSlice.actions