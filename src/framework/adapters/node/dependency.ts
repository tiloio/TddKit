import { sendDiscoveryDependenciesMessage } from "./communicate";
import { currentTestId, isDiscoveryPhase, isTestPhase } from "./env"
import { activateTests } from "./test";
import { Dependency } from "./types";


export const TestDependencies = (id: string, ...dependencies: Dependency[]): Dependency => {
    if (isTestPhase) {
        if (id === currentTestId) activateTests();
        return;
    }
    if (!isDiscoveryPhase) return;

    const dependency: Dependency = {
        id,
        dependencies
    };

    if (dependencies.length > 0) sendDiscoveryDependenciesMessage(dependency);


    // todo check recursivly all dependent dependencies are created like in resources (maybe use the same method)
    return dependency
}