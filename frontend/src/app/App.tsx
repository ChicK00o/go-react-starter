import React from 'react';
import Pager from "./routing/Pager";
import Navbar from "./routing/Navbar";

const App: React.FC = () => {
    return (
        <div className="zeroMargin">
            <Navbar/>
            <div className="bodyCentering">
                <Pager/>
            </div>
        </div>
    )
};

export default App
