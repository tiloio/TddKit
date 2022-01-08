const {helloWorld} = require('./hello-world');
const assert = require('assert');

for(let i = 0; i < 100; i++) {
    test("returns name " + i, () => {

        const text = helloWorld('test framework');

        assert.equal(text, 'Hello test framework');
    });

    test("some failing test " + i, () => {

        const text = helloWorld('test framework');

        assert.equal(text, 'Hello test');
    });
}