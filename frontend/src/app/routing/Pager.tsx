import React from "react";
import universal from 'react-universal-component'
import {useSelector} from "react-redux";
import {RootState} from "../store";

const load = (props: any) => Promise.all([
    import(`./pages${props.page}`)
]).then(proms => proms[0]);

const UniversalComponent = universal(load, {
    chunkName: props => props.page,
    resolve: props => require.resolveWeak(`./pages${props.page}`)
});

const Pager = () => {
    const {type, routesMap} = useSelector(
        (state: RootState) => state.location
    );
    return (<UniversalComponent page={routesMap[type]}/>)
};

export default Pager
