import { Component} from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { Tweet } from '../model/tweet';
import { TweetService } from '../services/tweet.service';


@Component({
  selector: 'app-tweets',
  templateUrl: './tweets.component.html',
  styleUrls: ['./tweets.component.css']
})
export class TweetsComponent {

  constructor(private tweetService : TweetService) { 
    this.getTweets();
  }

  tweets!: Observable<Tweet[]>;

  getTweets(){
    this.tweets = this.tweetService.getTweets()
  }

}
