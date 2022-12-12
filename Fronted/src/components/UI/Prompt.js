import React, { useState, useRef } from 'react'
import { Button, CenterPopup } from 'antd-mobile'
import classes from './Prompt.module.css'
import { useAddMutation } from '../../store/Api/addPartener-api';


const AAA = (props) => {
    // const [visible, setVisible] = useState(true)
    const myNameInp = useRef();
    const otherNameInp = useRef();
    localStorage.setItem('partnerState',false) 

    const updata = props.updata
    const [addFn, { error: addError }] = useAddMutation()
    const submitHandler = (e) => {
        e.preventDefault();
        const myName = parseInt(myNameInp.current.value)
        const otherName = parseInt(otherNameInp.current.value)
        addFn({
            UserName: myName,
            partner_number: otherName
        }).then(res => {
            console.log(res);
            if (!res.error) {
                localStorage.setItem('partnerState',true)
                updata() 
                // const partner = stateSlice.getInitialState
                // const partnerState = partner(() => state)
                // console.log(partnerState);
            }
        })
    }




    return (
        <>
            <CenterPopup
                visible={props.visible}
                onMaskClick={() => {
                    // setVisible(false)
                    props.getstate(true)
                }}
            >
                <div className={classes.div}>
                    <form onSubmit={submitHandler}>
                        <input
                            type='number'
                            ref={myNameInp}
                            className={classes.my}
                            placeholder='请输入自己的账号'></input>
                        <input
                            ref={otherNameInp}
                            type='number'
                            className={classes.other}
                            placeholder='请输入舍友的账号'></input>
                        <button
                            className={classes.btn}>提交</button>
                    </form>
                </div>
            </CenterPopup>
        </>
    )
}

export default AAA