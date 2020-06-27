import React, {FormEvent, useContext, useEffect, useState} from 'react';
import {useDispatch, useSelector} from "react-redux";
import {RootState} from "../../app/rootReducer";
import ReactJson from "react-json-view";
import {closeBackend, pingBackend} from "./jsonSlice";
import {ChangeHandler} from "../../utilities/types";
import {WebSocketContext} from "../../websocket/websocketProvider";

const ResponseJson = () => {
    const dispatch = useDispatch();

    const {value} = useSelector(
        (state: RootState) => state.jsonHolder
    );

    const [intervalRef, setIntervalRef] = useState<NodeJS.Timeout | undefined>(undefined);

    useEffect(() => {
        return () => {
            if (intervalRef !== undefined) {
                clearInterval(intervalRef);
            }
        }
    }, [intervalRef]);

    const onPingServer = () => {
        if (intervalRef === undefined) {
            setIntervalRef(setInterval(() => {
                dispatch(pingBackend());
            }, 1000))
        }
    };

    const onCloseServer = () => {
        onStopRefresh();
        dispatch(closeBackend());
    };


    const onStopRefresh = () => {
        if (intervalRef !== undefined) {
            clearInterval(intervalRef);
            setIntervalRef(undefined)
        }
    };

    const [data, setData] = useState("not set");
    const handleChange: ChangeHandler = event => {
        const {value} = event.target;
        setData(value);
    };

    const ws = useContext(WebSocketContext);
    const onFormSubmit = (event: FormEvent) => {
        event.preventDefault();
        ws?.sendMessage(data)
    };

    return (
        <div>
            <div>
                <ReactJson src={value} theme="solarized"/>
                <button onClick={onPingServer}>Start</button>
                <button onClick={onStopRefresh}>Stop Refresh</button>
            </div>
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

export default ResponseJson;
