import {helloWorld} from './hello-world';
import assert from 'assert';

test("returns name", () => {
    const text = helloWorld('test framework');

    assert.equal(text, 'Hello test framework');
});

test("some failing test", () => {

    const text = helloWorld('test framework');

    assert.equal(text, 'Hello test');
});