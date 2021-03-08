import { createStore, combineReducers } from "redux";
import { AuthAction } from "./pages/login/action";
import loginReducer from "./pages/login/reducer";
import { AuthState } from "./pages/login/state";
import { Action } from "redux";

export interface ReduxState {
  authState: AuthState;
}

export type ReduxAction = AuthAction | Action;

export default createStore(
  combineReducers({
    loginReducer,
  })
);
