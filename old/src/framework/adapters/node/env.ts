import { WRONG_PHASE } from "./errors";
import { activateTests } from "./test";
import { ResourceCreated } from "./types";

const {
    TEST_FRAMEWORK_PHASE,
    TEST_FRAMEWORK_ID,
    TEST_FRAMEWORK_TEST_ID,
    TEST_FRAMEWORK_RESOURCE_ID,
    TEST_FRAMEWORK_RESOURCES
} = process.env;

// console.log("ENV", {TEST_FRAMEWORK_PHASE, TEST_FRAMEWORK_ID, TEST_FRAMEWORK_TEST_ID})
const allowedPhases = [
    'DISCOVERY',
    'CREATE',
    'TEST_RUN',
];

export const phase = (TEST_FRAMEWORK_PHASE ?? 'TEST_RUN') as 'TEST_RUN' | 'DISCOVERY' | 'CREATE';
export const runId = TEST_FRAMEWORK_ID ?? 'no id given';

if (allowedPhases.every(allowedPhase => allowedPhase !== phase)) WRONG_PHASE(phase, allowedPhases.join(', '));

export const isDiscoveryPhase = TEST_FRAMEWORK_PHASE === 'DISCOVERY';
export const isCreatePhase = TEST_FRAMEWORK_PHASE === 'CREATE';
export const isTestPhase = TEST_FRAMEWORK_PHASE === 'TEST_RUN';

export const currentTestId = TEST_FRAMEWORK_TEST_ID || (isTestPhase && activateTests());

const throwNoResourceIdError = () => {
    throw new Error('Phase CREATE needs the environment variable TEST_FRAMEWORK_RESOURCE_ID with a resource id!');
}
const throwNoResourceError = () => {
    throw new Error('Phase TEST needs the environment variable TEST_FRAMEWORK_RESOURCES with the resources array as stringified JSON!');
}

export const currentResourceId = TEST_FRAMEWORK_RESOURCE_ID || (isCreatePhase && throwNoResourceIdError());
const resources = TEST_FRAMEWORK_RESOURCES || (isTestPhase && throwNoResourceError());

export const currentTestRunResources = (): ResourceCreated[] => JSON.parse(resources);