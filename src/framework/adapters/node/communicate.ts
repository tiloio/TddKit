import { Dependency, DiscoveredResource, DiscoveryType, Resource, TestSuiteType } from "./types";


const LOG_PREFIX = "__TFW:"

export const sendMessage = (message: string) => console.log(LOG_PREFIX + message);

export const sendObjectMessage = (obj: any) => sendMessage(JSON.stringify(obj));

export const sendDiscoveryTestSuiteMessage = (suite: TestSuiteType) => sendObjectMessage({
    type: DiscoveryType.TestSuite,
    id: suite.id,
    dependencies: suite.dependencies,
    resources: suite.resources.map(toDiscoveredResource)
});

export const sendDiscoveryTestMessage = (name: string) => sendObjectMessage({
    type: DiscoveryType.Test,
    name,
});


const toDiscoveredResource = (resource: Resource) => ({
    id: resource.id,
    resources: resource.resources.map(toDiscoveredResource)
});