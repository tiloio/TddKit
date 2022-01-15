const express = require('express')
const { countAll, addItem } = require('./repository');

const app = express()
const port = 3000

app.get('/', async (req, res) => {
    res.send(`Got ${await countAll()}`);
});

app.post('/', async (req, res) => {
    await addItem();
    res.send('item added');
});

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});