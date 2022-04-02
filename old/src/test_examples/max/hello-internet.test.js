const {helloInternet} = require('./hello-internet');
const assert = require('assert');

for(let i = 0; i < 100; i++) {
    test('returns internet name ' + i, () => {

        const text = helloInternet('test framework');

        assert.equal(text, 'Hello test framework internet');
    });
}
