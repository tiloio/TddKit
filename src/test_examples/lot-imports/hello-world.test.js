const {helloWorld} = require('./hello-world');
const dep1 = require('./hello-world copy');
const dep2 = require('./hello-world copy 2');
const dep3 = require('./hello-world copy 3');
const dep4 = require('./hello-world copy 4');
const dep5 = require('./hello-world copy 5');
const dep6 = require('./hello-world copy 6');
const dep7 = require('./hello-world copy 7');
const dep8 = require('./hello-world copy 8');
const dep9 = require('./hello-world copy 9');
const dep10 = require('./hello-world copy 10');
const dep11 = require('./hello-world copy 11');
const dep12 = require('./hello-world copy 12');
const assert = require('assert');

for(let i = 0; i < 100; i++) {
    test("returns name " + i, () => {

        const text = helloWorld('test framework');

        assert.equal(text, 'Hello test framework');
    });

    test("some failing test " + i, () => {

        const text = helloWorld('test framework');

        assert.equal(text, 'Hello test' + 
        dep1.someDep +
         dep2.someDep +
         dep3.someDep +
         dep4.someDep +
         dep5.someDep +
         dep6.someDep +
         dep7.someDep +
         dep8.someDep +
         dep9.someDep +
         dep10.someDep +
         dep11.someDep +
         dep12.someDep 
         );
    });
}