const sqlite3 = require('sqlite3').verbose();
const path = require('path');

const dbPath = path.resolve(__dirname, 'database.sqlite');

const db = new sqlite3.Database(dbPath, (err) => {
  if (err) {
    console.error('Error opening database ' + err.message);
  } else {
    console.log('Connected to the SQLite database.');
  }
});

const initializeDatabase = () => {
  db.serialize(() => {
    db.run(`CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      email TEXT NOT NULL UNIQUE
    )`, (err) => {
      if (err) {
        console.error('Error creating table ' + err.message);
      }
    });
  });
};

const addUser = (name, email) => {
  return new Promise((resolve, reject) => {
    db.run(`INSERT INTO users (name, email) VALUES (?, ?)`, [name, email], function(err) {
      if (err) {
        reject(err);
      } else {
        resolve({ id: this.lastID, name, email });
      }
    });
  });
};

const getUsers = () => {
  return new Promise((resolve, reject) => {
    db.all(`SELECT * FROM users`, [], (err, rows) => {
      if (err) {
        reject(err);
      } else {
        resolve(rows);
      }
    });
  });
};

const closeDatabase = () => {
  db.close((err) => {
    if (err) {
      console.error('Error closing database ' + err.message);
    } else {
      console.log('Database connection closed.');
    }
  });
};

module.exports = {
  initializeDatabase,
  addUser,
  getUsers,
  closeDatabase
};