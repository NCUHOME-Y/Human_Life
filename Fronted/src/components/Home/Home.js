import React, { useCallback, useState } from "react";
import MainMenu from "../Menu/mainMenu";
import classes from './Home.module.css'
import horn from '../assests/horn.png'
import moment from "moment/moment";
import { useSelector } from "react-redux";
import sleep from '../assests/sleep.png'
import dog from '../../image/dog.jpg'
import heng from '../assests/+.png'
import shu from '../assests/-.png'
import Prompt from '../UI/Prompt.js'


const Login = () => {

  const [visible, setVisible] = useState(false)
  const state = localStorage.getItem('exsit')  
  const [,setUpdata] = useState('')
  const forceUpdate = useCallback(() => {console.log('111'); setUpdata({})}, []);
  const partnerState = localStorage.getItem('partnerState')
  const partnerStateTwo = localStorage.getItem('partnerStateTwo')
  const partnerStateThree = localStorage.getItem('partnerStateThree') 
  console.log(partnerState);
  let hour = moment().hours()
  let minute = moment().minutes()
  if (hour < 10) {
    hour = '0' + hour
  }
  if (minute < 10) {
    minute = '0' + minute
  }

  const getstate = (state2) => {
    setVisible(!state2)
  }


  return (
    <>
      <div className={classes.top}>
        <div className={classes.a}>
          <div className={classes.time}>
            {`${hour}:${minute}`}
            <div className={classes.word}>
              现在是午休时间,请保持安静
            </div>
          </div>
        </div>

        <div className={classes.horn}>
          <img src={horn}></img>
          <p>去你妈的不关灯傻逼傻逼傻逼傻逼傻逼傻逼傻逼</p>
        </div>

        <div className={classes.inputMessage}>
          <input placeholder="请输入一条公告"></input>
          <button><div className={classes.s}>发送</div></button>
        </div>

      </div>

      <div className={classes.hidden}>
      </div>
      <hr></hr>
      <div className={classes.gridContainer}>
        <div className={classes.gridItemOne}>
          <img src={sleep} className={classes.sleep}></img>
          <div className={classes.state}>{state}</div>
          <div className={classes.user}>
            <img src={dog}></img>
            <div>小家园</div>
          </div>
        </div>
        <div
          className={classes.gridItemTwo}
        >
          { partnerState
            ? <>
              <img
                src={heng}
                onClick={() => { setVisible(true) }}
                className={classes.heng}></img>
              <img src={shu} className={classes.shu}></img>
              {visible && <Prompt updata={forceUpdate} getstate={getstate} visible={visible}></Prompt>}
            </>
            : <>
              <div className={classes.b}>
              <img src={sleep} className={classes.sleep}></img>
              <div className={classes.state}>{state}</div>
              <div className={classes.user}>
                <img src={dog}></img>
                <div>小家园</div>
              </div></div> 
            </>
          }

        </div>
        <div className={classes.gridItemThree}>
          { !partnerStateTwo 
          ? <>
          <img
            src={heng}
            onClick={() => { setVisible(true) }}
            className={classes.heng}></img>
          <img src={shu} className={classes.shu}></img>
          {visible && <Prompt getstate={getstate} visible={visible}></Prompt>}
        </>
        : <>
          <div className={classes.b}>
          <img src={sleep} className={classes.sleep}></img>
          <div className={classes.state}>{state}</div>
          <div className={classes.user}>
            <img src={dog}></img>
            <div>小家园</div>
          </div></div> 
        </>
          }

        </div>
        <div className={classes.gridItemFour}>

          { !partnerStateThree 
            ? <>
            <img
              src={heng}
              onClick={() => { setVisible(true) }}
              className={classes.heng}></img>
            <img src={shu} className={classes.shu}></img>
            {visible && <Prompt getstate={getstate} visible={visible}></Prompt>}
          </>
          : <>
            <div className={classes.b}>
            <img src={sleep} className={classes.sleep}></img>
            <div className={classes.state}>{state}</div>
            <div className={classes.user}>
              <img src={dog}></img>
              <div>小家园</div>
            </div></div> 
          </>
          }

        </div>
      </div>
      <MainMenu></MainMenu>
    </>
  )
}


export default Login