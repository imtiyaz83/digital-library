const express = require("express");
const app = express();
const port = 3000;

app.get("/", (req, res) => {
  res.send("User Service is up!");
});

app.listen(port, () => {
  console.log(`User service running at http://localhost:${port}`);
});
