import * as React from "react";
import { useDispatch, useSelector } from "react-redux";
import { ReduxState } from "../../store";
import { Login } from "./login";
import { AuthAction, loginAction } from "./action";

export class ActionDispatcher {
  constructor(private dispatch: React.Dispatch<AuthAction>) {}

  public login = (authInfo: {}): void => {
    this.dispatch(loginAction(authInfo));
  };
}

const authSelector = (state: ReduxState) => state.authState;

export default () => {
  const authState = useSelector(authSelector);
  return (
    <Login value={authState} action={new ActionDispatcher(useDispatch())} />
  );
};
