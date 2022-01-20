import { sendDiscoveryTestMessage, sendObjectMessage } from "./communicate";
import { isDiscoveryPhase } from "./env";

let testActive = false;

export const test = async (name, fn) => {
    if (isDiscoveryPhase) sendDiscoveryTestMessage(name);
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