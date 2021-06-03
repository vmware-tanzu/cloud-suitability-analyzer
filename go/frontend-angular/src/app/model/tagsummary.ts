export class TagSummary {
    name: string;
    count: number;

    constructor(name: string, count: number) {
        this.name = name;
        this.count = count;
    }
}

export class TagsSummary {
    tags: TagSummary [];
    constructor(tags: TagSummary[]) {
        this.tags = tags;
    }
}