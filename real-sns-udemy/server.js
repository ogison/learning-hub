const express = require("express");
const app = express();
const userRoute = require("./routes/users");
const authRoute = require("./routes/auth");
const postRoute = require("./routes/post");
const PORT = 3000;
const mongoose = require("mongoose");
require("dotenv").config();

// connect databse
mongoose
  .connect(process.env.MONGOURL)
  .then(() => {
    console.log("DBと接続中");
  })
  .catch((err) => {
    console.error(err);
  });

// ミドルウェア
app.use(express.json());
app.use("/api/users", userRoute);
app.use("/api/auth", authRoute);
app.use("/api/posts", postRoute);

app.listen(PORT, () => console.log("サーバーが起動しました"));
