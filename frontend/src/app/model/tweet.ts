export class Tweet {

    tweetId!: number;
    text!: string;
    username!: string;

    constructor(tweetId: number, text: string, username: string) {
        this.tweetId = tweetId;
        this.text = text;
        this.username = username;
    }
}