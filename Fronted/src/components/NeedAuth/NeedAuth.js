import { useSelector } from "react-redux"
import { Navigate, useLocation } from "react-router-dom"
const NeedAuth = props => {
    const auth = useSelector(state => state.auth)
    console.log(auth);
    const location = useLocation()
    console.log(location);
    return (auth.islogged
            ?props.children
            // :<Navigate 
            // to={'/sleepquiet'} 
            // replace
            // state={{preLocation:location}} ></Navigate>
            :<Navigate 
            to={'/sleepquiet'} 
            ></Navigate>
        )
}
export default NeedAuth