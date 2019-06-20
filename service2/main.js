const express = require('express');

const server = express();
const HOST_NAME = '0.0.0.0';
const PORT = 3000;

server.get('/', (req, res) => {
    json = {
        time: new Date(),
        message: "hello world",
        service: "app2"
    };
    res.json(json);
})

server.listen(PORT, HOST_NAME, ()=> {
    console.log("express server is running");
    console.log(`port: ${PORT}`);
})