import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { Tweet } from '../model/tweet';
import { TweetService } from '../services/tweet.service';


@Component({
  selector: 'app-tweets',
  templateUrl: './tweets.component.html',
  styleUrls: ['./tweets.component.css']
})
export class TweetsComponent implements OnInit {

  constructor(private tweetService : TweetService) { 
    this.getTweets();
  }

  ngOnInit(): void {
  }

  tweets!: Observable<Tweet[]>;

  getTweets(){
    this.tweets = this.tweetService.getTweets()
  }

}
