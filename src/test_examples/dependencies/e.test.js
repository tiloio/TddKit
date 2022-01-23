const { TestSuite, test } = require('test-framework');
const { dTest } = require('./d.test');

exports.eTest = TestSuite("e", {dependencies: [dTest]});

test("e test", () => {
    return true;
});
