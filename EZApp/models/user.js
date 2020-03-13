var mongoose = require("mongoose");
const Schema = mongoose.Schema;

var UserSchema = new Schema(
  {
    // required = must input
    uid: { type: Number, required: [true, "missing1"] },
    username: { type: String, required: "missing2" },
    createTime: { type: Date, required: "missing3" },
    items: { type: Array, required: "missing4" }
  },
  // version key: each time update __v will update too
  { versionkey: false }
);

// get virtual attribute
UserSchema.virtual("id").get(function() {
  return this._id.toHexString();
});

// when use toJSON to convert, it will contain virtuals attribute
UserSchema.set("toJSON", {
  virtuals: true
});

module.exports = mongoose.model("User", UserSchema);
