import * as React from "react";
import { HashRouter, Switch, Route, Redirect } from "react-router-dom";
import Layout from "./components/layout/layout";
import { SubPage } from "./subpage";
import Login from "./pages/login/container";
import _ from "lodash";
export const AppRoute = () => {
  // todo: get from redux
  let isAuthorized = true;
  return (
    <>
      <HashRouter>
        <Switch>
          {/* todo: /以外のrouteに対してもloginチェック */}
          <Route exact={true} path="/">
            {isAuthorized ? (
              <Redirect to="/app/dashboard" />
            ) : (
              <Redirect to="/login" />
            )}
          </Route>
          <PrivateRoute path="/app" component={Layout} />
          <Route path="/sub" component={SubPage} />
          <Route path="/login" component={Login} />
        </Switch>
      </HashRouter>
      ,
    </>
  );
};

const PrivateRoute = (props: any) => {
  const rest = _.omit(props, ["component"]);
  return <Route {...rest} render={(innerProps) => <Route {...props} />} />;
};
