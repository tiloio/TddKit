const { TestSuite, test } = require('test-framework');

exports.aTest = TestSuite("a");

test("a test", () => {
    return true;
});
