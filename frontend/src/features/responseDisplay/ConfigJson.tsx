import React, {useContext} from 'react';
import {useSelector} from "react-redux";
import {RootState} from "../../app/rootReducer";
import ReactJson from "react-json-view";
import {WebSocketContext} from "../../websocket/websocketProvider";
import log from "loglevel";

const ConfigJson = () => {

    const {config} = useSelector(
        (state: RootState) => state.jsonHolder
    );

    const ws = useContext(WebSocketContext);

    return (
        <div>
            <ReactJson
                src={config}
                theme="solarized"
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

export default ConfigJson;
