const express = require('express')
const { MongoClient } = require('mongodb')

const app = express()
const port = 3000

const uri = "mongodb://root:rootpassword@localhost:27017";


const execCollection = async (actionFn) => {
    const client = new MongoClient(uri);
    await client.connect();
    const database = client.db("itemsDb");
    const collection = database.collection("someItems");

    try {
        return await actionFn(collection);
    } finally {
        await client.close();
    }
}

async function addItem() {
    await execCollection(async (collection) => {

        const result = await collection.insertOne({
            date: Date.now()
        });
        console.log(`A document was inserted with the _id: ${result.insertedId}`);
    });
}
async function countAll() {
    return await execCollection(async (collection) => {
        return await collection.countDocuments();
    });
}

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