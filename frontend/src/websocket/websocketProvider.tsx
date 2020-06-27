import React, { createContext, PropsWithChildren, ReactNode } from 'react'
import { useDispatch } from 'react-redux';
import { jsonResponse } from 'features/responseDisplay/jsonSlice';
import * as log from 'loglevel';

const WebSocketContext = createContext<Websocket | null>(null);

export { WebSocketContext }

interface Websocket {
    webSocket : WebSocket,
    sendMessage : (message: string) => void
}

const WebSocketProvider = ({children} : PropsWithChildren<ReactNode>) => {
    let webSocket : WebSocket;
    let ws : Websocket;

    const dispatch = useDispatch();

    const sendMessage = (message : string) => {
        const payload = {
            type: "sendMessage",
            data: message
        };
        webSocket.send(JSON.stringify(payload))
        // dispatch(updateChatLog(payload));
    };


    webSocket = new WebSocket("ws://127.0.0.1:5000/api/ws");
    webSocket.onmessage = ev => {
        const payload = JSON.parse(ev.data);
        dispatch(jsonResponse(payload));
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
