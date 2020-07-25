import {Action, ThunkDispatch} from "@reduxjs/toolkit";
import {RootState} from "../app/store";
import {Message} from "./websocketMessageHandler";
import log from "loglevel";
// import {DisplayResponse, displayResponse} from "../features/customFeatures/.../displaySlice";

export const customMessageHandler = (payload: Message, dispatch: ThunkDispatch<RootState, unknown, Action<string>>) => {
    switch (payload.type) {
        case "display": {
            log.debug(payload.type + " : " + payload.body);
            // const value = payload.body as DisplayResponse;
            // dispatch(displayResponse(value));
            break;
        }
    }
};
