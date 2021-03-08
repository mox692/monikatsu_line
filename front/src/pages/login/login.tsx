import * as React from "react";
import { Link } from "react-router-dom";
import { ActionDispatcher } from "./container";

import { AuthState } from "./state";
// todo: impl
const login = () => {};

interface Props {
  value: AuthState;
  action: ActionDispatcher;
}

export const Login = (props: Props) => {
  console.log("props");
  console.log(props);
  return (
    <>
      <h2>this is login page !</h2>
      <form>
        mail
        <input type="text"></input>
        pass
        <input type="text"></input>
        <button onClick={() => props.action.login({})}></button>
        <p>isAuthorizedd: {props.value}</p>
      </form>
    </>
  );
};
