import React, {FormEvent, useContext, useState} from 'react';
import {useDispatch} from "react-redux";
import {closeBackend} from "../../../features/responseDisplay/jsonSlice";
import {InputChangeHandler} from "../../../utilities/types";
import {WebSocketContext} from "../../../websocket/websocketProvider";

const System = () => {
    const dispatch = useDispatch();

    const onCloseServer = () => {
        dispatch(closeBackend());
    };

    const [data, setData] = useState("not set");
    const handleChange: InputChangeHandler = event => {
        const {value} = event.target;
        setData(value);
    };

    const ws = useContext(WebSocketContext);
    const onFormSubmit = (event: FormEvent) => {
        event.preventDefault();
        ws?.sendMessage("user", data)
    };

    return (
        <div>
            <form onSubmit={onFormSubmit}>
                <label>Message Backend</label>
                <input value={data} onChange={handleChange}/>
                <button type="submit">Send over websocket</button>
            </form>
            <div>
                <label>To Close Backend server</label>
                <button onClick={onCloseServer}>Close</button>
            </div>
        </div>
    );
};

export default System;
