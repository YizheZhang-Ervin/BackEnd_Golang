// express is a function
let express = require("express");
let promiseFS = require("./promiseFS");
let bodyParser = require("body-parser");

// create express web service (listener function) use returned app to operate service
let app = express();

// convert client request to object and save in request.body by using middleware
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(bodyParser.raw());

//static resources file request handling -> search in given directory
app.use(express.static("./static"));

// When not find given resources -> 404 handling / redirect
app.use((request, response, next) => {
  response.status(404);
  response.send("not found");
  // response.redirect(301, "http://...");
});

// data API handling(get/delete/put/head)
// get
app.get("/a", function(request, response, next) {
  // // get by promise
  // let { lx = "pro" } = request.query;
  // promiseFS
  //   .readFile("./package.json")
  //   .then(result => {
  //     result.JSON.parse(result);
  //     result = lx === "dev" ? result.devDependancies : result.devDependancies;
  //     response.status(200);
  //     response.type("application/json");
  //     response.send(result);
  //   })
  //   .catch(err => {
  //     response.status(500);
  //     response.type("application/json");
  //     response.send(err);
  //   });
  response.send("hello a");
});

// post: after receive information, server save it to local xx.json
app.post("/b", function(request, response, next) {
  // // complex method:
  // let chunk = "";
  // // receive information bu spliting several times
  // response.on("data", chart => {
  //   chunk += chart;
  // });
  // response.on("end", () => {
  //   // qs.stringify & qs.parse <=> JSON & urlencoded
  //   let qs = require("qs");
  //   console.log(qs.parse(chunk));
  // });
  // simple method:
  console.log(request.body);
});

// listen port + start server
app.listen(3000, () => {
  console.log("start server");
});
