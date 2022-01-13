export { default as expect } from 'expect';
import { sendMessage } from './communicate';

export const test = async (name, fn) => {
    try {
        await fn();
        sendMessage(JSON.stringify({
            name,
        }));
    } catch (err) {
        sendMessage(JSON.stringify({
            name,
            err: JSON.stringify({
                message: err.message,
            })
        }));
    }
}