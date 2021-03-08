import * as React from "react";
import * as ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { AppRoute } from "./routes";
import store from "./store";
import { ThemeProvider } from "@material-ui/styles";
import GlobalTheme from "./theme/index";

ReactDOM.render(
  <Provider store={store}>
    <ThemeProvider theme={GlobalTheme.default}>
      <AppRoute />
    </ThemeProvider>
  </Provider>,
  document.getElementById("root")
);
