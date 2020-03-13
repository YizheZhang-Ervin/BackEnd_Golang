"use strict";

module.exports = app => {
  const models = require("./models/index");
  const routes = require("./routes/index");
  routes(app);
};
