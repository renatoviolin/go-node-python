const express = require('express')
const connectDB = require('./db/db')
const { helloWorld, getAll, singleJson } = require('./controller/payload')

// =================================
const cluster = require('cluster');
const port = 3000
const app = express()
const numWorkers = require('os').cpus().length;

// =================================
connectDB()
app.use(express.urlencoded({ extended: false }))
app.use(express.json())

// =================================
app.get('/hello', helloWorld)
app.get('/single_json', singleJson)
app.get('/mongo_json', getAll)

if (cluster.isMaster) {
  console.log('Master cluster setting up ' + numWorkers + ' workers...');
  for (var i = 0; i < numWorkers; i++) {
    cluster.fork();
  }

  cluster.on('online', function (worker) {
    console.log('Worker ' + worker.process.pid + ' is online');
  });

  cluster.on('exit', function (worker, code, signal) {
    console.log('Worker ' + worker.process.pid + ' died with code: ' + code + ', and signal: ' + signal);
    console.log('Starting a new worker');
    cluster.fork();
  });
} else {
  // Run the server!
  app.listen(port, function () {
    console.log('Process ' + process.pid + ' is listening to all incoming requests');
  });
}