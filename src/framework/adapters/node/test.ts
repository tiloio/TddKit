import { sendObjectMessage } from "./communicate";

let testActive = false;

export const test = async (name, fn) => {
    if (!testActive) return;

    try {
        await fn();
        sendObjectMessage({ name });
    } catch (err) {
        sendObjectMessage({
            name,
            err: JSON.stringify({
                message: err.message,
            })
        });
    }
}

export const activateTests = () => testActive = true;