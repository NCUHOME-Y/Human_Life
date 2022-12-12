import React  from "react";
import { Link, Navigate, NavLink, useNavigate} from 'react-router-dom'
import {useDispatch, useSelector} from 'react-redux'
import { logout } from "../../store/reducer/authSlice";
import {
  AppOutline,
  UserOutline,
} from 'antd-mobile-icons'
import { TabBar } from 'antd-mobile'
import classes from './mainMenu.module.css'
import home from '../assests/home.png'
import about from '../assests/about.png'

const MainMenu = () => {
    const UserName = localStorage.getItem('user')
    const navigate = useNavigate()
    const auth = useSelector(state => state.auth)
    console.log(auth);
    localStorage.setItem('isLogged',auth.islogged)
    const dispatch = useDispatch()
  
    return(
      <>
      <div>
        <div className={classes.bottom}>
          <div className={classes.bottomLeft}>
            <Link to={`/sleepquiet/user/${UserName}/home`}>
            <div className={classes.a}>
              <img src={home}></img>
              <div className={classes.word}>主页</div>
            </div>
            </Link>
          </div>
          <div className={classes.bottomRight}>
            <Link to={`/sleepquiet/user/${UserName}/profile`}>
            <div className={classes.b}>
            <img src={about}></img>
            <div className={classes.word}> 我的</div>
            </div>
            </Link>
          </div>
        </div>
      </div>
      {/* <div className={classes.menu}>
        <TabBar onChange={e => navigate(e) }>
          {tabs.map(item => (
            <TabBar.Item key={item.key} icon={item.icon} title={item.title} />
          ))}
        </TabBar>
        </div> */}
        {/* <ul>
        <li>
          <Link to={'/home'}>首页</Link>
        </li>

        {
          !auth.islogged && <li>
          <Link to={'/'}>登录/注册</Link>
          </li>
        }

        {
          auth.islogged &&
          <> 
          <li>
          <Link to={'/profile'}>Id</Link>
          </li>
          <li>
          <Link 
          to={'/'}
          onClick={ () => dispatch(logout())}>退出登录</Link>
          </li>
          </>
        }
     
      </ul> */}
      </>
    )
}
export default MainMenu