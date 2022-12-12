import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import {BrowserRouter as Router} from "react-router-dom";
import { Provider } from 'react-redux';
import store from './store';
import {} from './index.css'

//设置移动端状态
document.documentElement.style.fontSize = 20000 / 788 +'vw'
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <Provider store={store}>
         <Router>
            <App/>
        </Router>
    </Provider>
   
);

