const {random} = require('./random');
const assert = require('assert');

for(let i = 0; i < 1000; i++) {
    test('randomness ' + i, () => {
        assert(random());
    });
}
