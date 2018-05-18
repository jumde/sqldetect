'use strict';

const express = require('express');
const mysql = require('mysql');

const connection = mysql.createConnection({
  host: 'db',
  user: 'root',
  password: 'secret',
  database: 'db',
});
connection.connect((err) => {
  if (err) {
    console.log("Error");
    throw err;
  }
  console.log('Connected!');
});

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
  res.send('Pranjal\n');
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
