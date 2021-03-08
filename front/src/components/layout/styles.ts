import {
  createStyles,
  makeStyles,
  useTheme,
  Theme,
} from "@material-ui/core/styles";

export default makeStyles((theme: Theme) => ({
  root: {},
  wrap: {
    color: "pink",
    backgroundColor: "#fafaff",
    display: "flex",
  },
  content: {
      flexGrow: 1,
    padding: theme.spacing(3),
  },
  right: {
    width: "100%",
  },
  left: {},
}));
