const mongoose = require('mongoose')

module.exports = mongoose.model('Payload', mongoose.Schema({},{strict: false}), 'sample')

