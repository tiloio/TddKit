const { TestDependencies: TestDependency, test } = require('test-framework');
const { aTest } = require('./a.test');

exports.dTest = TestDependency("d", aTest);

test("a test", () => {
    throw new Error('FAILED!');
});
