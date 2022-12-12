import {configureStore} from '@reduxjs/toolkit'
import { authApi } from './Api/auth-api'
import {setupListeners} from "@reduxjs/toolkit/query";
import { authSlice } from './reducer/authSlice';
import { addApi } from './Api/addPartener-api';
const store = configureStore ({
    reducer:{
        [authApi.reducerPath] : authApi.reducer,
        [addApi.reducerPath] : addApi.reducer,
        auth : authSlice.reducer
  },
    middleware:(getDefaultMiddleware) => 
        getDefaultMiddleware().concat(
        authApi.middleware,
        addApi.middleware
        )  
});

setupListeners(store.dispatch)

export default store