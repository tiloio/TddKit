const { TestDependencies: TestDependency, test } = require('test-framework');

exports.aTest = TestDependency("a");

test("a test", () => {
    return true;
});
