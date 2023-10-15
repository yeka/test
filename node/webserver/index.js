const express = require('express')
const web = express()

// == Admin specific routings, protected by admin specific middleware
const admin = express.Router()
admin.use((req, res, next) => {
    console.log('--admin only')
    res.locals.admin = 'the admin' // passing data from middleware to controller
    next()
})

admin.get('/', (req, res) => {
    console.log('----admin get: admin is', res.locals.admin)
    res.json('ok')
})
admin.get('/update', (req, res) => {
    console.log('----admin update: admin is', res.locals.admin)
    res.json('ok')
})

// == Common routings, use logging middleware
web.use((req, res, next) => {
    console.log('incoming request')
    next()
})
web.use('/admin', admin) // set prefix for admin routings. logging middleware affects the admin routings.

web.get('/', (req, res) => {
    console.log('--root hello')
    res.json('ok')
})
web.get('/user', (req, res) => {
    console.log('--get user')
    res.json('ok')
})

web.listen(3001, '127.0.0.1', () => {
    console.log('server started')
})
