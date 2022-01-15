export type Dependency = {
    id: string,
    dependencies: Dependency[]
}

export type Resource = {
    id: string,
    create: () => Promise<void>,
    init: () => Promise<void>,
    resources: Resource[]
}

export type DiscoveredResource = {
    id: string,
    resources: DiscoveredResource[]
}

export enum DiscoveryType  {
    dependency = 'DEPENDENCY',
    resource = 'RESOURCE'
}