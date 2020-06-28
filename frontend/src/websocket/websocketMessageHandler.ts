import {Action, ThunkDispatch} from "@reduxjs/toolkit";
import {RootState} from "../app/rootReducer";
import {jsonResponse} from "../features/responseDisplay/jsonSlice";
import log from "loglevel";

interface Message {
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
        switch (payload.type) {
            case "display": {
                log.debug(payload.type + " : " + payload.body);
                // const value = payload.body as DisplayResponse;
                // dispatch(displayResponse(value));
                break;
            }
            case "ping_pong":
            case "system":
            default: {
                log.debug(payload.type + " : " + payload.body);
                break;
            }
        }
        dispatch(jsonResponse(payload));
    } else {
        dispatch(jsonResponse(rawPayload));
    }
};
