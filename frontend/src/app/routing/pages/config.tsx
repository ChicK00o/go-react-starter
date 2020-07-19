import React, {useContext} from 'react';
import {useSelector} from "react-redux";
import ReactJson from "react-json-view";
import log from "loglevel";
import {RootState} from "../../store";
import {WebSocketContext} from "../../../websocket/websocketProvider";

const Config = () => {

    const {config} = useSelector(
        (state: RootState) => state.jsonHolder
    );

    const ws = useContext(WebSocketContext);

    return (
        <div>
            <p>App Config</p>
            <ReactJson
                src={config}
                onEdit={edit => {
                    if ((typeof edit.existing_value) !== (typeof edit.new_value)) {
                        log.warn(edit);
                        return false
                    } else {
                        log.warn(typeof edit.existing_value);
                        log.warn(typeof edit.new_value);
                        log.warn(edit);
                        ws?.sendMessage("config", edit.updated_src);
                    }
                }}
            />
        </div>
    );
};

export default Config;
