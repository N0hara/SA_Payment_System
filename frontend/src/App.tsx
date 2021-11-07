import React, { useEffect} from "react";
import { BrowserRouter as Router, Switch, Route} from "react-router-dom";
import {
  createStyles,
  makeStyles,
  Theme,
} from "@material-ui/core/styles";

import CssBaseline from "@material-ui/core/CssBaseline";
import Home from "./components/Home";
import Bills from "./components/Bills";
import BillCreate from "./components/BillCreate";
import SignIn from "./components/Signin";
import Navbar from "./components/Navbar";

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: "flex",
    },
    nav : {background : "#239B56"},
    title: {
      flexGrow: 1,
    },
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      marginLeft: drawerWidth,
      width: `calc(100% - ${drawerWidth}px)`,
      transition: theme.transitions.create(["width", "margin"], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: 36,
    },
    hide: {
      display: "none",
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
      whiteSpace: "nowrap",
    },
    drawerOpen: {
      width: drawerWidth,
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    drawerClose: {
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      overflowX: "hidden",
      width: theme.spacing(7) + 1,
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9) + 1,
      },
    },
    toolbar: {
      display: "flex",
      alignItems: "center",
      justifyContent: "flex-end",
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
    },
    a: {
      textDecoration: "none",
      color: "inherit",
    },
  })
);

export default function MiniDrawer() {
  const classes = useStyles();
  const [token, setToken] = React.useState<String>("");
  
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);
  
  if (!token) {
    return <SignIn />;
  }
  

 
  return (
    <div className={classes.root}>
      <Router>
        <CssBaseline />
        {token && (
          <>
          <Navbar />
          <main className={classes.content}>
          <div className={classes.toolbar} />
          <div>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/bills" component={Bills} />
              <Route exact path="/bill/create" component={BillCreate} />
            </Switch>
          </div>
          </main>
        </>
        )}

        
      </Router>
    </div>
  );
}