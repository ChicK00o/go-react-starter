import React, {useContext} from 'react';
import {useSelector} from "react-redux";
import ReactJson from "react-json-view";
import {RootState} from "../../store";
import {WebSocketContext} from "../../../websocket/websocketProvider";

const Response = () => {
    // const dispatch = useDispatch();

    const {value} = useSelector(
        (state: RootState) => state.jsonHolder
    );

    const ws = useContext(WebSocketContext);
    const onPingServer = () => {
        // dispatch(pingBackend());
        ws?.sendMessage("data", {})
    };

    return (
        <div>
            <p>API Response for reading</p>
            <div>
                <ReactJson src={value}/>
                <button onClick={onPingServer}>Get Data</button>
            </div>
        </div>
    );
};

export default Response;
