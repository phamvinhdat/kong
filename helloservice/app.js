const express = require('express')
const app = express()

app.get('/', function (req, res) {
    console.log("have req")
    res.send('hello world')
})

app.listen(9600)