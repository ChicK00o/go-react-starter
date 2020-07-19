import React from 'react';
import Navbar from "./routing/Navbar";
import Pager from "./routing/Pager";

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
