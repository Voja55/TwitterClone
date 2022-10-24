export class Tweet {

    id!: number;
    text!: string;
    title!: string;

    constructor(id: number, text: string, title: string) {
        this.id = id;
        this.text = text;
        this.title = title;
    }
}