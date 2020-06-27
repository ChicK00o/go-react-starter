import { combineReducers } from "@reduxjs/toolkit"
import jsonHolderReducer from 'features/responseDisplay/jsonSlice'

const rootReducer = combineReducers({
    jsonHolder: jsonHolderReducer
});

export type RootState = ReturnType<typeof rootReducer>

export default rootReducer
