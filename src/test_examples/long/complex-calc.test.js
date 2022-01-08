const {complexCalc} = require('./complex-calc');
const assert = require('assert');

for(let i = 0; i < 50; i++) {
    test('calcs cool ' + i, () => {

        const result = complexCalc(3, 2);

        assert.equal(result, 100);
    });

    test('calcs hot ' + i, () => {

        const result = complexCalc(3, 2);

        assert.notEqual(result, 0);
    });
}