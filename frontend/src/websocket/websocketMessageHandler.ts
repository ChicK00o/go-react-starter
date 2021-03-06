import {Action, ThunkDispatch} from "@reduxjs/toolkit";
import {jsonResponse, configResponse} from "../features/responseDisplay/jsonSlice";
import log from "loglevel";
import {RootState} from "../app/store";
import {customMessageHandler} from "./customMessageHandler";

export interface Message {
    body: any,
    type: string,
}

export const sendMessageHandler = (websocket: WebSocket, type: string, message: any) => {
    let data: Message = {
        type: type,
        body: message
    };
    websocket.send(JSON.stringify(data))
};

export const receivedMessageHandler = (data: any, dispatch: ThunkDispatch<RootState, unknown, Action<string>>) => {
    const rawPayload = JSON.parse(data);
    const payload = rawPayload as Message;
    if (payload !== undefined) {
        customMessageHandler(payload, dispatch);
        switch (payload.type) {
            case "display": {
                dispatch(jsonResponse(payload));
                break;
            }
            case "config": {
                dispatch(configResponse(payload.body));
                break;
            }
            case "ping_pong":
            case "system":
            default: {
                log.debug(payload.type + " : " + payload.body);
                dispatch(jsonResponse(payload));
                break;
            }
        }
    } else {
        dispatch(jsonResponse(rawPayload));
    }
};
