const { test, TestSuite } = require('test-framework');
const { repositoryTest } = require('./repository.test'); // TODO problematic
const { default: fetch } = require("node-fetch");
const assert = require('assert');

const webServerSavesOneObjectTest = TestSuite('webServerSavesOneObject', { dependencies: [repositoryTest] });

test("saves one object", async () => {
    const response = await fetch("http://localhost:3000", { method: 'POST' });

    assert.equal(await response.text(), "item added");
});
