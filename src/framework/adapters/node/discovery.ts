import { sendDiscoveryTestSuiteMessage } from "./communicate";
import { currentTestId, isDiscoveryPhase, isTestPhase } from "./env"
import { activateTests } from "./test";
import { TestSuiteCreateOptions, TestSuiteType } from "./types";

export const TestSuite = (id: string, options?: TestSuiteCreateOptions): TestSuiteType => {
    if (isTestPhase) {
        if (id === currentTestId) activateTests();
        return;
    }
    if (!isDiscoveryPhase) return;

    const suite: TestSuiteType = {
        id,
        dependencies: options?.dependencies ?? [],
        resources: options?.resources ?? [],
    };
    sendDiscoveryTestSuiteMessage(suite);

    return suite
}