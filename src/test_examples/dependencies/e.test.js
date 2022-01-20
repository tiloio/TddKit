const { TestDependencies: TestDependency, test } = require('test-framework');
const { dTest } = require('./d.test');

exports.eTest = TestDependency("e", dTest);

test("a test", () => {
    return true;
});
