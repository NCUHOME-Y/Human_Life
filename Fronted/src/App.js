import React from "react";
import {Routes,Route} from 'react-router-dom'
import NeedAuth from "./components/NeedAuth/NeedAuth";
import AuthFormPage from "./Pages/AuthFormPages";
import HomePage from "./Pages/HomePage";
import Clean from "./Pages/Profile/Clean/clean";
import MyMessage from "./Pages/Profile/MyMessage/myMessage";
import MyPartener from "./Pages/Profile/myPartener/myPartener";
import ProfilePage from "./Pages/ProfilePage";
function App() {

 //const auth = useSelector(state => state.auth)
  
  const UserName = localStorage.getItem('user')
  return (
    //<Layout>
    <Routes>
            <Route  path={`/sleepquiet/user/${UserName}/home`} element={<NeedAuth><HomePage/></NeedAuth>}/>
            <Route  path={`/sleepquiet/user/${UserName}/profile`} element={<NeedAuth><ProfilePage/></NeedAuth>}/>
            {/* <Route  path='/auth-form' element={<AuthFormPage/>}></Route> */}
            <Route  path='/sleepquiet' element={<AuthFormPage/>}></Route>
            <Route  path={`/sleepquiet/user/${UserName}/profile/myMessage`} element={<MyMessage></MyMessage>}></Route>
            <Route  path={`/sleepquiet/user/${UserName}/profile/clean`} element={<Clean></Clean>}></Route>
            <Route  path={`/sleepquiet/user/${UserName}/profile/myPartener`} element={<MyPartener></MyPartener>}></Route>
            {/* <Route  path='/sleep' element={<MyMessage></MyMessage>}></Route> */}
    </Routes>
    //</Layout>
  );
}

export default App;
