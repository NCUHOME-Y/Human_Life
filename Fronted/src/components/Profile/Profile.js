import React, { useCallback, useEffect, useState } from "react";
import MainMenu from "../Menu/mainMenu";
import classes from './Profile.module.css'
import dog from '../../image/dog.jpg'
import { Link } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import Declarative from "../UI/Declarative";
import { Button } from "antd-mobile";

const Profile = () => {
    // const state = useSelector( state => state.state)
    // localStorage.setItem('exsit','我在寝')
    localStorage.setItem('isSet',false)
    const location = window.location.pathname + window.location.search
    console.log(location);
    const getValue = (e) => {
        // dispatch(setState(e.target.value))
        // console.log(state);
        localStorage.setItem('exsit',e.target.value)
        localStorage.setItem('isSet',true)
        console.log(exit);
    }
    const isSet = localStorage.getItem('isSet')
    const exit = localStorage.getItem('exsit')
    return(
        <>
        <div className={classes.body}>
           <div className={classes.top}>
             <img src={dog}></img>
             <div>小家园</div>
           </div>
            <div className={classes.father}>
                <Link to={`${location}/myMessage`}>
                    <div className={classes.myMessage}><span>账号信息</span></div>
                </Link>
                <Link to={`${location}/clean`}>
                    <div className={classes.clean}><span>值日信息</span></div>
                </Link>
               
                <div className={classes.exist}>
                    <span>状态</span>
                    {
                        isSet && <select onChange={getValue}>
                        <option>我在寝</option>
                        <option>外出中</option>
                        <option>睡眠中</option>
                    </select>    
                    }
                    {
                         (!isSet && exit == '我在寝' ) && <select 
                         
                         onChange={getValue}>
                        <option selected='selected'>我在寝</option>
                        <option>外出中</option>
                        <option>睡眠中</option>
                    </select>    
                    }
                     {
                         (!isSet && exit == '外出中' ) && <select 
                         onChange={getValue}>
                        <option selected='selected'>外出中</option>
                        <option>我在寝</option>
                        <option>睡眠中</option>
                    </select>    
                    }
                     {
                         (!isSet && exit == '睡眠中' ) && <select                       
                         onChange={getValue}>
                        <option selected='selected'>睡眠中</option>
                        <option>我在寝</option>
                        <option>外出中</option>
                        
                    </select>    
                    }
                </div>

                {/* <Link onClick={<Declarative></Declarative>}> */}
                    <div className={classes.myBedfriend}><span>我的舍友</span><Declarative></Declarative></div>
                {/* </Link> */}
                <div className={classes.setTimemessage}><span>设置时间段文字</span><Declarative></Declarative></div>
                <div className={classes.setting}><span>设置</span><Declarative></Declarative></div>    
           </div>
        </div>
        <MainMenu></MainMenu>
        </>
    )
}

export default Profile