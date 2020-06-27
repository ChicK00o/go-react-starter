import React from 'react';
import ResponseJson from "../features/responseDisplay/ResponseJson";

const App: React.FC = () => {
    return (
        <div className="container">
            <div className="container">
                <h1>Response</h1>
                <p>API Response for reading</p>
                <ResponseJson />
            </div>
        </div>
    )
};

export default App
