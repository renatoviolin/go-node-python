const express = require('express')
const connectDB = require('./db/db')
const { helloWorld, getAll, singleJson } = require('./controller/payload')
const cluster = require("cluster");
const os = require('os');
const clusterWorkerSize = os.cpus().length;


// =================================
connectDB()


// =================================
const fastify = require('fastify')({
  logger: false,
  disableRequestLogging: true
});

fastify.get('/hello', helloWorld)
fastify.get('/single_json', singleJson)
fastify.get('/mongo_json', getAll)


// Run the server!
const start = async () => {
  try {
    await fastify.listen(3000);
    console.log(`server listening on ${fastify.server.address().port} and worker ${process.pid}`);

  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
}

if (clusterWorkerSize > 1) {
  if (cluster.isMaster) {
    for (let i = 0; i < clusterWorkerSize; i++) {
      cluster.fork();
    }

    cluster.on("exit", function (worker) {
      console.log("Worker", worker.id, " has exited.")
    })
  } else {
    start();
  }
} else {
  start();
}