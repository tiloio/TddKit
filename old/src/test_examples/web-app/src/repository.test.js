const { test, TestSuite } = require('test-framework');
const { addItem, countAll } = require('./repository');
const assert = require('assert');

module.exports.repositoryTest = TestSuite('repository');

test("add item adds one item", async () => {
    await addItem();
    await addItem();
    await addItem();

    assert.equal(await countAll(), 3);
});

