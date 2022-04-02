const {helloWorld} = require('./hello-world');
const assert = require('assert');

/** comments i like
 */
test("returns name", () => {
    //! a comment
    const text = helloWorld('test framework');

    assert.equal(text, 'Hello test framework');
});

test("some failing test", () => {

    const text = helloWorld('test framework');

    assert.equal(text, 'Hello test');
});