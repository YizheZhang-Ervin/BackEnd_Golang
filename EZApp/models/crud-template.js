var mongoose = require("mongoose");
require("./model.js");

var Book = mongoose.model("Book");

// insert
var book = new Book({
  name: "123",
  author: "abc",
  publishTime: new Date()
});

book.author = "def";
book.save(function(err) {
  console.log("insert status", err ? "failed" : "success");
});

// retrieve all
Book.find({}, function(err, docs) {
  if (err) {
    console.log("retrieve status: failed");
    return;
  }
  console.log("result:", docs);
});

// retrieve one and update/delete
Book.findOne({ author: "def" }, function(err, doc) {
  if (err) {
    console.log("retrieve status: failed");
    return;
  }
  // retrieve and update
  doc.author = "efg";
  // retrieve and delete
  if (doc) {
    doc.remove();
  }
  doc.save();
  console.log("result:", doc);
});

// retrieve several conditions by $or/$and
var cond = {
  $or: [{ author: "efg" }, { author: "abc" }]
};
Book.find(cond, function(err, docs) {
  if (err) {
    console.log("retrieve status: failed");
    return;
  }
  console.log("condition:", cond, "result:", docs);
});
