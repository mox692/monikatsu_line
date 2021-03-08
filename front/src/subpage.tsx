import * as React from "react";
import { Link } from "react-router-dom";

export const SubPage = () => {
  return (
    <>
      <h2>this is home sub !</h2>
      <p>
        <Link to="/app">to homepage</Link>
      </p>
    </>
  );
};
