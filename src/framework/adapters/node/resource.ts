export type Resource = {
    id: string,
    create: () => Promise<void>,
    init: () => Promise<void>,
    dependencies: Resource[]
}

export const TestResources = async (resources: Resource[]) => {
    // todo check recursivly all dependent resources are created like in dependencies (maybe use the same method)
    return await Promise.all(resources.map(resource => resource.create()))
}