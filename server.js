// require libs
let express = require("express"),
  app = express(),
  port = process.env.xx_PORT || 3000,
  mongoose = require("mongoose"),
  bodyParser = require("body-parser"),
  config = require("./EZApp/config/config"),
  ejs = require("ejs"),
  path = require("path");

// self made utils
let { dataHandle } = require("./EZApp/utils/utils");

// connect to DB
mongoose.connect(config.mongodb);
mongoose.Promise = global.Promise;

// use template
app.set("views", path.join("./EZApp/", "views"));
app.set("view engine", "ejs");

// middleware 1: xx=yy convert to xx:yy
app.use(bodyParser.urlencoded({ extended: true }));

// Middleware 2: static handling
app.use(express.static("./static"));
// 404 handling
app.use((req, res, next) => {
  res.status(404);
  // res.send(`Request ${res.path} not exist`);
  res.render("index", { data: res.statusCode + " not found " + res.path });
});

// Middleware 3: header handling
app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header(
    "Access-Control-Allow-Headers",
    "Origin,X-Requested-With,Content-Type,Accept"
  );
  next();
});

// start
const initApp = require("./EZApp/app");
initApp(app);
app.listen(port, () => console.log(`start at ${port}`));
