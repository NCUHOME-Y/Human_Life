import {createApi, fetchBaseQuery} from "@reduxjs/toolkit/dist/query/react"
export const authApi = createApi({
    reducerPath : 'authApi',
    baseQuery : fetchBaseQuery({
        baseUrl : 'http://43.143.227.115:5500/hack/'
    }),
    endpoints(build){  //指定该Api功能
        return{
            register:build.mutation({
                query(user){
                    console.log(user);
                    return {
                        url : 'register',
                        method : 'post',
                        body : user
                    }
                }
            }),
            login:build.mutation({
              query(user){
                console.log(user);
                return{
                    url : 'login',
                    method : 'post',
                    body : user
                }
              }  
            }) 
        }
    }
}) 

export const {
    useRegisterMutation,
    useLoginMutation
} = authApi

