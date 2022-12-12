import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/dist/query/react";

const token = localStorage.getItem('token')

export const addApi = createApi({
    reducerPath: 'addApi',
    baseQuery: fetchBaseQuery({
        baseUrl: 'http://43.143.227.115:5500/Logined/'
    }),
    endpoints(build) {
        return {
            add: build.mutation({
                query(user) {
                    console.log(user);
                    return {
                        url: 'addPartner',
                        method: 'post',
                        headers: {
                            'token': token
                        },
                        body: user
                    }
                }
            })
        }
    }
})

export const {
    useAddMutation
} = addApi