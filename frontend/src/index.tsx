import React from 'react'
import ReactDOM from 'react-dom'
import App from './app/App'
import './index.css'
import store from "./app/store";
import {Provider} from 'react-redux'
import WebSocketProvider from "./websocket/websocketProvider";

const render = () => {
    ReactDOM.render(
        <Provider store={store}>
            <WebSocketProvider>
                <App/>
            </WebSocketProvider>
        </Provider>,
        document.getElementById('root')
    )
};

render();
