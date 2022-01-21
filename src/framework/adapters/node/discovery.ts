import { sendDiscoveryDependenciesMessage, sendDiscoveryResourcesMessage } from "./communicate";
import { currentTestId, isDiscoveryPhase, isTestPhase } from "./env"
import { activateTests } from "./test";
import { Dependency, Resource } from "./types";

export type TestSuiteOptions = {
    dependencies?: Dependency[],
    resources?: Resource[]
}
export const TestSuite = (id: string, options?: TestSuiteOptions): Dependency => {
    if (isTestPhase) {
        if (id === currentTestId) activateTests();
        return;
    }
    if (!isDiscoveryPhase) return;

    const dependency: Dependency = {
        id,
        dependencies: options?.dependencies ?? []
    };

    sendDiscoveryDependenciesMessage(dependency);
    if (options?.resources?.length > 0) sendDiscoveryResourcesMessage(options.resources);


    // todo check recursivly all dependent dependencies are created like in resources (maybe use the same method)
    return dependency
}