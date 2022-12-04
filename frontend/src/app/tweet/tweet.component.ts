import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Tweet } from '../model/tweet';
import { TweetLikes } from '../model/tweetLikes';
import { StoreService } from '../services/store-service.service';
import { TweetService } from '../services/tweet.service';

@Component({
  selector: 'app-tweet',
  templateUrl: './tweet.component.html',
  styleUrls: ['./tweet.component.css']
})
export class TweetComponent implements OnInit {

  constructor(public store : StoreService, private tweetService : TweetService, private router : Router) {
    
  }

  ngOnInit(): void {
    this.getLikes()
  }

  @Input()
  tweet! : Tweet

  getLikes() {
    this.tweetService.getLikes(this.tweet).subscribe(data => {
      console.log(data);
      this.tweet.likes = data.likes;
    })
  }

  likeTweet(){
    this.tweet.username = this.store.getUsername();
    this.tweetService.likeTweet(this.tweet).subscribe(data => {
      console.log(data);
      this.getLikes()
    })
  }

  redirect() {
    this.router.navigateByUrl("/login");
  }

}
