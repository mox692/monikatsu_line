import * as React from "react";
import { Link } from "react-router-dom";
import { Route, Switch, withRouter } from "react-router";
import { Dashboard } from "../../pages/dashboard/dashboard";
import Header from "../header/header";
import Sidebar from "../sidebar/sidebar";
import useStyles from "./styles";

// todo: propsはany型でいいのか？
const Layout = (props: any) => {
  let classes = useStyles();
  return (
    <>
      <div className={classes.root}>
        <div className={classes.wrap}>
          <div className={classes.left}>
            <Sidebar />
          </div>
          <div className={classes.right}>
            <Header />
            <div className={classes.content}>
              <Switch>
                <Route path="/app/dashboard" component={Dashboard} />
              </Switch>
              <h2>this is layout page !</h2>
              <p>
                <Link to="/sub">to subpageas</Link>
              </p>
              <p>
                <button onClick={() => console.log(props)}>see props?</button>
              </p>
              <p>
                <button onClick={() => props.history.goBack()}>
                  do props.history.back ?
                </button>
              </p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default withRouter(Layout);
