import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { of, Subscriber } from 'rxjs';
import { Tweet } from '../model/tweet';
import { map} from 'rxjs';


@Component({
  selector: 'app-tweets',
  templateUrl: './tweets.component.html',
  styleUrls: ['./tweets.component.css']
})
export class TweetsComponent implements OnInit {

  constructor() { 
    this.getTweets();
  }

  ngOnInit(): void {
  }

  tweets: Array<Tweet> = new Array;

  getTweets(){
    this.tweets[0] = new Tweet(1, "text1", "title1");
    this.tweets[1] = new Tweet(2, "text2", "title2");
    this.tweets[2] = new Tweet(3, "text3", "title3")
  }

}
