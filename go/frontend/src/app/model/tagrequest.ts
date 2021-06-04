export class TagRequest {
    name: string;
    selected: boolean;

    constructor(name: string, selected: boolean) {
        this.name = name;
        this.selected = selected;
    }
}

export class TagRequests {
    tags: TagRequest [];
    constructor(tags: TagRequest[]) {
        this.tags = tags;
    }
}