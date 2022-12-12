import React, { useState } from 'react'
import { Button, Modal } from 'antd-mobile'
import classes from './Declarative.module.css'
const Declarative = () => {
    const [visible, setVisible] = useState(false)
    return (
      <>
        <Button
          className={classes.btn}
          block
          onClick={() => {
            setVisible(true)
          }}
        >
          点下试试
        </Button>
        <Modal
          visible={visible}
          content='前端跑路啦'
          closeOnAction
          onClose={() => {
            setVisible(false)
          }}
          actions={[
            {
              key: 'confirm',
              text: '滚',
            },
          ]}
        />
      </>
    )
  }
  export default Declarative