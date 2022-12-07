import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { Tweet } from '../model/tweet';
import { User } from '../model/user';
import { StoreService } from '../services/store-service.service';
import { TweetService } from '../services/tweet.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

  constructor(public store : StoreService, private userService : UserService, private tweetService : TweetService) { }

  ngOnInit(): void {
    this.getTweets()
  }

  user : User = new User();
  statusDisplayName : boolean = false;
  statusDescription : any;
  statusPassword : any;
  changeDescription(){}
  changeDisplayName(){}
  changePassword(){}

  confirm() {
    console.log(this.store.getUsername());
    this.user.username = this.store.getUsername();
    let codeField = document.getElementById("code") as HTMLInputElement;
    this.userService.confirmAuth(this.user.username, +codeField.value).subscribe(data => {
      console.log(data);
    })
  }

  tweets!: Observable<Tweet[]>;

  getTweets(){
    this.tweets = this.tweetService.getTweets()
  }

  
}
