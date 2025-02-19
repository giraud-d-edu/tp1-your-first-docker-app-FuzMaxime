// filepath: /c:/Users/maxim/Documents/Docker/test/my-js-server-1/src/index.js
const express = require("express");
const path = require("path");
const database = require("./database");

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware to parse JSON and URL-encoded data
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Serve static files from the public directory
app.use(express.static(path.join(__dirname, "public")));

// Initialize the database
database.initializeDatabase();

// Route to handle form submission
app.post("/submit", (req, res) => {
  const { name, email } = req.body;
  database
    .addUser(name, email)
    .then(() => res.redirect("/success.html"))
    .catch((err) => res.status(500).send(err.message));
});

// Route to get all users
app.get("/users", (req, res) => {
  database
    .getUsers()
    .then((users) => res.json(users))
    .catch((err) => res.status(500).send(err.message));
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
