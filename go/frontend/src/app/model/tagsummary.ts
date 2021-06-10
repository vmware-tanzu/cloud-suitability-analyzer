export class TagSummary {
    name: string;
    count: number;
    tagData: any;

    constructor(name: string, count: number, tagData: any) {
        this.name = name;
        this.count = count;
        this.tagData = tagData;
    }
}

export class TagsSummary {
    tags: TagSummary [];
    constructor(tags: TagSummary[]) {
        this.tags = tags;
    }
}