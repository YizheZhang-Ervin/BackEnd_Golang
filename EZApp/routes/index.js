"use strict";

const userRoute = require("./user-route");

module.exports = app => {
  userRoute(app);
  // total first lay app url
  // app.use("/user", require("./xx"));
};
