const { TestSuite, test } = require('test-framework');
const { aTest } = require('./a.test');

exports.dTest = TestSuite("d", {dependencies: [aTest]});

test("d test", () => {
    throw new Error('FAILED!');
});
