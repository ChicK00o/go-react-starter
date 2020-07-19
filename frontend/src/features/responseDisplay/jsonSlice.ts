import {Action, createSlice, PayloadAction, ThunkDispatch} from "@reduxjs/toolkit";
import {AppThunk, RootState} from "../../app/store";
import axios from 'axios';
import * as log from 'loglevel';
import {PORT} from "../../constants";

interface JsonHolder {
    value: any
    config: any
}

const initialState: JsonHolder = {
    value: {},
    config: {},
};

const jsonHolder = createSlice({
    name: 'jsonHolder',
    initialState,
    reducers: {
        jsonResponse(state, {payload}: PayloadAction<any>) {
            log.debug(payload);
            state.value = payload;
        },
        configResponse(state, {payload}: PayloadAction<any>) {
            log.debug(payload);
            state.config = payload;
        },
    }
});

export const {
    jsonResponse,
    configResponse,
} = jsonHolder.actions;

export default jsonHolder.reducer

export const pingBackend = ():
    AppThunk => async dispatch => {
    await doThis(getPing, dispatch);
};

interface PingResponse {
    payload: any
}

async function getPing() {
    const url = "http://127.0.0.1:" + PORT + "/api/data";
    const {data} = await axios.get<PingResponse>(url);
    return data
}

export const closeBackend = ():
    AppThunk => async dispatch => {
    await doThis(getClose, dispatch);
};

async function getClose() {
    const url = "http://127.0.0.1:" + PORT + "/api/close";
    const {data} = await axios.get<PingResponse>(url);
    return data
}

async function doThis(getApi: () => Promise<any>, dispatch: ThunkDispatch<RootState, unknown, Action<string>>) {
    try {
        const data = await getApi();
        dispatch(jsonResponse(data));
    } catch (err) {
        if (err.response) {
            dispatch(jsonResponse(err.response));
        } else if (err.request) {
            dispatch(jsonResponse(err.request));
        } else {
            dispatch(jsonResponse(err));
        }
    }
}
