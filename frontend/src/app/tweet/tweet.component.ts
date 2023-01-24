import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Tweet } from '../model/tweet';
import { StoreService } from '../services/store-service.service';
import { TweetService } from '../services/tweet.service';

@Component({
  selector: 'app-tweet',
  templateUrl: './tweet.component.html',
  styleUrls: ['./tweet.component.css']
})
export class TweetComponent implements OnInit {

  constructor(private modalService: NgbModal, public store : StoreService, private tweetService : TweetService, private router : Router) {
    
  }

  ngOnInit(): void {
    this.getLikes()
  }

  @Input()
  tweet! : Tweet

  users! : string[];

  getLikes() {
    this.tweetService.getLikes(this.tweet).subscribe(data => {
      console.log(data);
      this.tweet.likes = data.likes;
    })
  }

  likeTweet(){
    //this.tweet.username = this.store.getUsername();
    this.tweetService.likeTweet(this.store.getUsername(), this.tweet.tweetId).subscribe(data => {
      console.log(data);
      this.getLikes()
    })
  }

  redirect() {
    this.router.navigateByUrl("/login");
  }

  openLg(content : any) {
    this.getLikeUsers()
    this.modalService.open(content, { size: 'lg' });

  }

  getLikeUsers() {
    this.tweetService.getLikeUsers(this.tweet.tweetId).subscribe(data => {
      console.log(data)
      this.users = data
    })
  }

}
