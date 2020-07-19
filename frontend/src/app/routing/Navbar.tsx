import React from "react";
import {NavLink} from 'redux-first-router-link'
import RoutesMap from "./RoutesMap";

const Navbar = () => {
    return (
        <nav>
            <ul>
                <li><NavLink to={RoutesMap.HOME}>Home</NavLink></li>
                <li><NavLink to={RoutesMap.CONFIG}>Config</NavLink></li>
                <li><NavLink to={RoutesMap.RESPONSE}>Response</NavLink></li>
                <li className="float-right"><NavLink to={RoutesMap.SYSTEM}>System</NavLink></li>
            </ul>
        </nav>
    )
};

export default Navbar
