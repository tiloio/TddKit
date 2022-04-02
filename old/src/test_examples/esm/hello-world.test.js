import {helloWorld} from './hello-world.js';
import assert from 'assert';

for(let i = 0; i < 500; i++) {
    test('returns name', () => {

        const text = helloWorld('test framework');

        assert.equal(text, 'Hello test framework');
    });

    test("some failing test", () => {

        const text = helloWorld('test framework');

        assert.equal(text, 'Hello test');
    });
}