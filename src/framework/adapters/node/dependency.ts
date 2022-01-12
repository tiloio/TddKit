export type Dependency = {
    id: string,
    dependencies: Dependency[]
}

export const TestDependencies = (id: string, ...dependencies: Dependency[]): Dependency => {
    // todo check recursivly all dependent dependencies are created like in resources (maybe use the same method)
    return {
        id, // maybe change this to the hand of the user
        dependencies
    }
}