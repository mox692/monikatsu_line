import { AuthState, initialState } from "./state";
import { AuthAction, ActionNames } from "./action";

export default (state: AuthState = initialState, action: AuthAction) => {
  switch (action.type) {
    case ActionNames.LOGIN:
      // todo: ここにfetch的なlogicを入れる
      console.log("called");
      return Object.assign({}, state, { isAuthorized: true });
    default:
      console.log("called!!");
      console.log(state);
      return state;
  }
};
