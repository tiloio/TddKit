const { TestDependencies: TestDependency, test } = require('test-framework');
const { bTest } = require('./b.test');

exports.cTest = TestDependency("c", bTest);

test("a test", () => {
    return true;
});
