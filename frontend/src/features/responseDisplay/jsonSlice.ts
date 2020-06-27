import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {AppThunk} from "../../app/store";
import axios from 'axios';
import * as log from 'loglevel';

interface JsonHolder {
    value: any
}


const initialState: JsonHolder = {
    value: {},
};

const jsonHolder = createSlice({
    name: 'jsonHolder',
    initialState,
    reducers: {
        jsonResponse(state, {payload}: PayloadAction<any>) {
            log.debug(payload);
            state.value = payload;
        },
    }
});

export const {
    jsonResponse,
} = jsonHolder.actions;

export default jsonHolder.reducer

export const pingBackend = ():
    AppThunk => async dispatch => {
    try {
        const data = await getPing();
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
};

interface PingResponse {
    payload: any
}

async function getPing() {
    const url = "http://127.0.0.1:5000/ping";
    const {data} = await axios.get<PingResponse>(url);
    return data
}

export const closeBackend = ():
    AppThunk => async dispatch => {
    try {
        const data = await getClose();
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
};

async function getClose() {
    const url = "http://127.0.0.1:5000/api/close";
    const {data} = await axios.get<PingResponse>(url);
    return data
}
