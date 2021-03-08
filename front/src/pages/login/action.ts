import { Action } from "redux";

export enum ActionNames {
  LOGIN = "login",
}

interface LoginAction extends Action {
  type: ActionNames.LOGIN;
  authInfo: {};
}

export const loginAction = (authInfo: {}): LoginAction => ({
  type: ActionNames.LOGIN,
  authInfo: authInfo,
});

export type AuthAction = LoginAction;
