const {helloInternet} = require('./hello-internet');
const assert = require('assert');

test('returns internet name', () => {

    const text = helloInternet('test framework');

    assert.equal(text, 'Hello test framework internet');
});