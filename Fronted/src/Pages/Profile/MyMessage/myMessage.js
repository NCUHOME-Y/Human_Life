import React from "react";
import classes from "./myMessage.module.css"
import dog from '../../../image/dog.jpg'
import { Link } from "react-router-dom";
import { useDispatch } from "react-redux";
import { logout } from "../../../store/reducer/authSlice";
const MyMessage = () => {
    const dispatch = useDispatch()
    const UserName = localStorage.getItem('user')

    return(
        <>
            <div className={classes.id}>
                <img src={dog}></img>
                <div>小家园</div>
            </div>

            <div className={classes.UserName}>
                <span>账号:{UserName}</span>
            </div>
            <div className={classes.resetPwd}>
                <span>修改密码</span>
            </div>

            <div className={classes.exit}>
                    <Link 
                    to={'/sleepquiet'}
                    replace    >
                        <button onClick={ () => dispatch(logout())}>退出登录</button>
                    </Link>
            </div>
        </>
    )
}
export default MyMessage