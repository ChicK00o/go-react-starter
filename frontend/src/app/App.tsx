import React from 'react';
import ResponseJson from "../features/responseDisplay/ResponseJson";
import ConfigJson from "../features/responseDisplay/ConfigJson";

const App: React.FC = () => {
    return (
        <div className="container">
            <div className="container">
                <h1>Response</h1>
                <p>API Response for reading</p>
                <ResponseJson />
                <p>App Config</p>
                <ConfigJson />
            </div>
        </div>
    )
};

export default App
