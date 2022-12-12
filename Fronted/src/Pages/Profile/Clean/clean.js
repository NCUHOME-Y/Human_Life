import React from "react";
import classes from './clean.module.css'
const Clean = () => {

    return (
        <>
            <div>
                <div className={classes.follow}>
                    <span>值日顺序</span>
                </div>
                <div className={classes.secondChild}>
                    <span>  1      小  家  园 </span>
                </div>
                <div className={classes.thirdChild}>
                    <span>  2      小  家  园 </span>
                </div>
                <div className={classes.forthChild}>
                    <span>  3      小  家  园 </span>
                </div>
                <div className={classes.fifthChild}>
                    <span>  4      小  家  园 </span>
                </div>
            </div>
            <div className={classes.day}>
                <span>值日天数</span>
            </div>
            <div className={classes.gridcontainer}>
                <div className={classes.griditem}>
                    <input type='checkbox'></input>
                       <span>周一</span>
                </div>
                <div className={classes.griditem}>
                <input type='checkbox'></input>
                       <span>周二</span>
                </div>
                <div className={classes.griditem}>
                <input type='checkbox'></input>
                       <span>周三</span>
                </div>
                <div className={classes.griditem}>
                <input type='checkbox'></input>
                       <span>周四</span>
                </div>
                <div className={classes.griditem}>
                <input type='checkbox'></input>
                       <span>周五</span>
                </div>
                <div className={classes.griditem}>
                <input type='checkbox'></input>
                       <span>周六</span>
                </div>
                <div className={classes.griditem}>
                <input type='checkbox'></input>
                       <span>周日</span>
                </div>               
            </div>
            <div >
                <button className={classes.btn}>完成</button>
            </div>
        </>
    )
}
export default Clean