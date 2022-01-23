const { TestSuite, test } = require('test-framework');
const { bTest } = require('./b.test');

exports.cTest = TestSuite("c", {dependencies: [bTest]});

test("c test", () => {
    return true;
});
