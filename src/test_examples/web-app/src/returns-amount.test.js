const { test, TestSuite } = require('test-framework');
const { repositoryTest } = require('./repository.test'); // TODO problematic
const { default: fetch } = require("node-fetch");
const assert = require('assert');

const webServerReturnsAmountTest = TestSuite('webServerReturnsAmount', { dependencies: [repositoryTest] });

test("returns amount of saved objects", async () => {
    await fetch("http://localhost:3000", { method: 'POST' });
    await fetch("http://localhost:3000", { method: 'POST' });
    const response = await fetch("http://localhost:3000");

    assert.equal(await response.text(), "Got 2");
});
