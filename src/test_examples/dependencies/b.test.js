const { TestDependencies: TestDependency, test } = require('test-framework');
const { aTest } = require('./a.test');

exports.bTest = TestDependency("b", aTest);

test("a test", () => {
    return true;
});