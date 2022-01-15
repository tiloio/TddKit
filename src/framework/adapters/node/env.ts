import { WRONG_PHASE } from "./errors";
import { activateTests } from "./test";

const { TEST_FRAMEWORK_PHASE, TEST_FRAMEWORK_ID, TEST_FRAMEWORK_TEST_ID } = process.env;

// console.log("ENV", {TEST_FRAMEWORK_PHASE, TEST_FRAMEWORK_ID, TEST_FRAMEWORK_TEST_ID})
const allowedPhases = [
    'TEST_RUN',
    'DISCOVERY'
];

export const phase = (TEST_FRAMEWORK_PHASE ?? 'TEST_RUN') as 'TEST_RUN' | 'DISCOVERY';
export const runId = TEST_FRAMEWORK_ID ?? 'no id given';


if (allowedPhases.every(allowedPhase => allowedPhase !== phase)) WRONG_PHASE(phase, allowedPhases.join(', '));

export const isDiscoveryPhase = TEST_FRAMEWORK_PHASE ===  'DISCOVERY';
export const isTestPhase = TEST_FRAMEWORK_PHASE ===  'TEST_RUN';


export const currentTestId = TEST_FRAMEWORK_TEST_ID || (isTestPhase && activateTests());