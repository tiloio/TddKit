import { sendDiscoveryTestSuiteMessage, sendResourceCreatedMessage } from "./communicate";
import { currentResourceId, currentTestId, currentTestRunResources, isCreatePhase, isDiscoveryPhase, isTestPhase, runId } from "./env"
import { activateTests } from "./test";
import { ResourceOptions, TestSuiteCreateOptions, TestSuiteDiscoveryType, TestSuiteType } from "./types";

export const TestSuite = (id: string, options?: TestSuiteCreateOptions): TestSuiteType => {
    const dependencies = options?.dependencies ?? [];

    if (isCreatePhase || isDiscoveryPhase) {
        const discoverySuite: TestSuiteDiscoveryType = {
            id,
            dependencies: dependencies,
            resources: options?.resources ?? [],
        };
        if (isCreatePhase) createResource(discoverySuite);
        if (isDiscoveryPhase) sendDiscoveryTestSuiteMessage(discoverySuite);
    }

    const suite: TestSuiteType = {
        id,
        dependencies,
        resources: []
    }

    if (isTestPhase) {
        suite.resources = currentTestRunResources();
        activateTestsAfterSuiteWasCreated(id);
    }

    return suite;
}

const activateTestsAfterSuiteWasCreated = (suiteId: string) => {
    if (suiteId === currentTestId) activateTests();
}

let resourceFound = false;

const createResourceOptions = (): ResourceOptions => ({
    runId: parseInt(runId, 10)
})

const createResource = async (suite: TestSuiteDiscoveryType) => {
    if (resourceFound) return;

    const resource = suite.resources.find(resource => resource.id === currentResourceId);
    if (!resource) return;

    const resourceCreated = await resource.create(createResourceOptions());
    // TODO validate resourceCreated.data fits into env variabke or use a temp file instead...
    sendResourceCreatedMessage(resourceCreated);
    process.exit(); // EXIT to prevent double creation of resource
}