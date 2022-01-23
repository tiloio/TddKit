const { TestSuite, test } = require('test-framework');
const { aTest } = require('./a.test');

exports.bTest = TestSuite("b", {dependencies: [aTest]});

test("b test", () => {
    return true;
});
