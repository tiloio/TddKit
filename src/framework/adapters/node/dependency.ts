import { sendMessage } from "./communicate";
import { isDiscoveryPhase } from "./env"

export type Dependency = {
    id: string,
    dependencies: Dependency[]
}

export const TestDependencies = (id: string, ...dependencies: Dependency[]): Dependency => {
    if (!isDiscoveryPhase) return;

    const dependency = {
        id,
        dependencies
    };

    if (dependencies.length > 0) {
        sendMessage(JSON.stringify({
            type: "DEPENDENCY",
            ...dependency
        }));
    }


    // todo check recursivly all dependent dependencies are created like in resources (maybe use the same method)
    return dependency
}