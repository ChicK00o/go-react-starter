import React, {createContext, PropsWithChildren, ReactNode} from 'react'
import {useDispatch} from 'react-redux';
import * as log from 'loglevel';
import {receivedMessageHandler, sendMessageHandler} from "./websocketMessageHandler";

const WebSocketContext = createContext<Websocket | null>(null);

export {WebSocketContext}

interface Websocket {
    webSocket: WebSocket,
    sendMessage: (type: string, message: any) => void
}

const WebSocketProvider = ({children}: PropsWithChildren<ReactNode>) => {
    let webSocket: WebSocket;
    let ws: Websocket;

    const dispatch = useDispatch();

    const sendMessage = (type: string, message: any) => {
        sendMessageHandler(webSocket, type, message);
    };

    webSocket = new WebSocket("ws://127.0.0.1:4999/api/ws");
    webSocket.onmessage = ev => {
        receivedMessageHandler(ev.data, dispatch);
    };

    webSocket.onopen = () => {
        log.warn("Websocket is now open");
    };

    webSocket.onclose = event => {
        log.warn("Socket Closed Connection: ", event);
    };

    webSocket.onerror = error => {
        log.error("Socket Error: ", error);
    };

    ws = {
        webSocket: webSocket,
        sendMessage
    };


    return (
        <WebSocketContext.Provider value={ws}>
            {children}
        </WebSocketContext.Provider>
    )
};

export default WebSocketProvider;
