import React from "react";
import { useState, useRef } from "react";
import { useDispatch } from "react-redux";
import { useLoginMutation, useRegisterMutation } from '../../store/Api/auth-api'
import { login } from "../../store/reducer/authSlice";
import { useLocation, useNavigate } from 'react-router-dom'
import classes from "./AuthForm.module.css"
import { UserCircleOutline, LockOutline } from 'antd-mobile-icons'
const AuthForm = () => {


    // const [Msg,setMsg] = useState('')
    // const [PassWord,setPassWord] = useState('')
    // const [Email,setEmail] = useState('')
    // 
    const [isLoginForm, setisLoginForm] = useState(true)
    //引入注册的Api
    const [regFn, { error: regError }] = useRegisterMutation()
    const [loginFn, { error: loginError }] = useLoginMutation()
    const UserNameInp = useRef();
    const pwdInp = useRef();
    // const HomeCodeInp = useRef();
    // const IsAdminInp = useRef();


    //const emailInp = useRef();

    const dispatch = useDispatch()

    const navigate = useNavigate()

    const location = useLocation()
    const from = location.state?.preLocation?.pathname || '/'

    localStorage.setItem('partnerState', false)
    localStorage.setItem('partnerStateTwo', false)
    localStorage.setItem('partnerStateThree', false)

    const submitHandler = (e) => {
        e.preventDefault();

        // 获取用户输入的内容
        const UserName = parseInt(UserNameInp.current.value);
        const PassWord = pwdInp.current.value;

        let IsAdmin
        // 处理登录功能
        if (isLoginForm) {
            console.log('登录 -->', UserName, PassWord);
            loginFn({
                UserName: UserName,
                PassWord: PassWord,
                // HomeCode : HomeCode
            }).then(res => {
                console.log(res);
                if (!res.error) {
                    dispatch(login(
                        {
                            //存储（jwt，用户信息）
                            token: res.data.token,
                            user: UserName,
                            isloginform: isLoginForm
                        }
                    ))
                    //增加需求：跳转至之前的路径
                    //navigate(from,{replace:true})
                    navigate(`/sleepquiet/user/${UserName}/home`)
                }
            })
        } else {

            // const HomeCode = HomeCodeInp.current.value;
            // const IsAdminPro =  IsAdminInp.current.value;
            console.log("注册----->", UserName, PassWord);
            // if(IsAdminPro === '舍长'){
            //     IsAdmin = true
            // }else{
            //     IsAdmin = false
            // }
            //console.log('注册 -->', UserName, PassWord, email);
            regFn({
                UserName,
                PassWord,
                // HomeCode,
                // IsAdmin
                // email
            }).then(res => {
                console.log(res);
                if (!res.error) {
                    // 注册成功
                    setisLoginForm(true);
                }
            });
        }
    }

    return (
        <div className={classes.total}>
            {/* <div style={{color:'black'}}> */}
            {/* {regError && regError.data.error.details.errors.map((item,key)=>{    
                    return [                       
                            <p key={key}>{item.message}</p>                      
                    ]
                })}  */}

            {/* {loginError && loginError.data.error.message} */}
            {/* {loginError && loginError.data.error.details.errors.map((item,key)=>{    
                    return [                       
                            <p key={key}>{item.message}</p>                      
                    ]
                })} */}
            {/* </div> */}
            <div className={classes.title}>

                宿  静
                {/* {isLoginForm
                ?"登录"
                :"注册"} */}

            </div>
            <form onSubmit={submitHandler}>
                <div className={classes.top}>

                    <div className={classes.UserName}>
                        <UserCircleOutline />
                        <input
                            type='number'
                            placeholder={'账号'}
                            ref={UserNameInp}>
                        </input>
                    </div>

                    <div className={classes.PassWord}>
                        <LockOutline />
                        <input
                            type='password'
                            placeholder={'密码'}
                            ref={pwdInp}></input>
                    </div>




                    {/* {
                        !isLoginForm &&
                        <>
                            <div className={classes.HomeCode}>
                                <h2>寝室号</h2>
                                <input       
                                type='text'
                                 placeholder={'寝室号'} ref={HomeCodeInp}></input>
                            </div>

                            <div className={classes.IsAdmin}>
                                <h2>身份</h2>
                                <input 
                                type='text'
                                 placeholder={'请输入你的身份：舍长/舍员'} 
                                 ref={IsAdminInp}></input>
                            </div>
                        </>
                    } */}


                    <div className={classes.reg}>
                        <button className={classes.btn}>
                            {isLoginForm
                                ? "登    录"
                                : '注    册'}
                        </button>
                        <div className={classes.qo}>
                            <a href="#"
                                onClick={e => {
                                    e.preventDefault()
                                    setisLoginForm(prevState => !prevState)
                                }}>
                                {isLoginForm
                                    ? "没有账号?点击注册"
                                    : "已有账号?点击登录"}
                            </a></div>
                    </div>
                </div>
            </form>
        </div>
    )
}

export default AuthForm