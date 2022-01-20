import { Dependency, DiscoveredResource, DiscoveryType, Resource } from "./types";


const LOG_PREFIX = "__TFW:"

export const sendMessage = (message: string) => console.log(LOG_PREFIX + message);

export const sendObjectMessage = (obj: any) => sendMessage(JSON.stringify(obj));

export const sendDiscoveryDependenciesMessage = (dependency: Dependency) => sendObjectMessage({
    type: DiscoveryType.dependency,
    ...dependency
});

export const sendDiscoveryTestMessage = (name: string) => sendObjectMessage({
    type: DiscoveryType.test,
    name,
});


const toDiscoveredResource = (resource: Resource) => ({
    id: resource.id,
    resources: resource.resources.map(toDiscoveredResource)
});
export const sendDiscoveryResourcesMessage = (resources: Resource[]) => sendObjectMessage({
    type: DiscoveryType.resource,
    resources: resources.map(toDiscoveredResource)
});