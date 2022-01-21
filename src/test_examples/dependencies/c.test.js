const { TestSuite, test } = require('test-framework');
const { bTest } = require('./b.test');

exports.cTest = TestSuite("c", {dependencies: [bTest]});

test("a test", () => {
    return true;
});
