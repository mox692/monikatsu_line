"use strict";
var __createBinding =
  (this && this.__createBinding) ||
  (Object.create
    ? function (o, m, k, k2) {
        if (k2 === undefined) k2 = k;
        Object.defineProperty(o, k2, {
          enumerable: true,
          get: function () {
            return m[k];
          },
        });
      }
    : function (o, m, k, k2) {
        if (k2 === undefined) k2 = k;
        o[k2] = m[k];
      });
var __setModuleDefault =
  (this && this.__setModuleDefault) ||
  (Object.create
    ? function (o, v) {
        Object.defineProperty(o, "default", { enumerable: true, value: v });
      }
    : function (o, v) {
        o["default"] = v;
      });
var __importStar =
  (this && this.__importStar) ||
  function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null)
      for (var k in mod)
        if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k))
          __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
  };
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppRoute = void 0;
var React = __importStar(require("react"));
var react_router_dom_1 = require("react-router-dom");
var home_1 = require("./home");
var subpage_1 = require("./subpage");
exports.AppRoute = function () {
  return React.createElement(
    React.Fragment,
    null,
    React.createElement(
      react_router_dom_1.HashRouter,
      null,
      React.createElement(
        react_router_dom_1.Switch,
        null,
        React.createElement(react_router_dom_1.Route, {
          exact: true,
          path: "/",
          component: home_1.Home,
        }),
        React.createElement(react_router_dom_1.Route, {
          path: "/sub",
          component: subpage_1.SubPage,
        })
      )
    ),
    ","
  );
};