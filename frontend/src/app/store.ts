import {
    Action,
    combineReducers,
    configureStore,
    getDefaultMiddleware,
    ThunkAction
} from "@reduxjs/toolkit";

import { connectRoutes } from 'redux-first-router'
import RoutesMap from "./routing/RoutesMap";
import jsonHolderReducer from "../features/responseDisplay/jsonSlice";

const {
    middleware: routerMiddleware,
    enhancer: routerEnhancer,
    reducer: routerReducer,
} = connectRoutes(RoutesMap);

const rootReducer = combineReducers({
    jsonHolder: jsonHolderReducer,
    location: routerReducer,
});


const store = configureStore({
    reducer: rootReducer,
    enhancers: defaultEnhancers => [routerEnhancer, ...defaultEnhancers],
    middleware: [...getDefaultMiddleware({
        immutableCheck : {
            warnAfter: 1000
        },
        serializableCheck : {
            warnAfter: 1000
        }
    }), routerMiddleware]
});

export type AppDispatch =  typeof store.dispatch

export type RootState = ReturnType<typeof rootReducer>
export type AppThunk = ThunkAction<void, RootState, unknown, Action<string>>

export default store
