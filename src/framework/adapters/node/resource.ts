import { sendDiscoveryResourcesMessage } from "./communicate";
import { isDiscoveryPhase } from "./env";
import { Resource } from "./types";

export const TestResources = async (resources: Resource[]) => {
    if (!isDiscoveryPhase) return;


    if (resources.length > 0) sendDiscoveryResourcesMessage(resources);

    // todo check recursivly all dependent resources are created like in dependencies (maybe use the same method)
    return await Promise.all(resources.map(resource => resource.create()))
}