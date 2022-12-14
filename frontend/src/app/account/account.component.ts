import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/internal/Observable';
import { Profile } from '../model/profile';
import { Tweet } from '../model/tweet';
import { User } from '../model/user';
import { ProfileService } from '../services/profile.service';
import { StoreService } from '../services/store-service.service';
import { TweetService } from '../services/tweet.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

  constructor(public store : StoreService, private userService : UserService, private tweetService : TweetService, private profileService : ProfileService, private router: Router) { }

  ngOnInit(): void {
    this.getTweets()
    this.getProfile()
  }

  user : User = new User();
  profile : Profile = new Profile();

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
    this.tweets = this.tweetService.getTweetsByUser(this.store.getUsername())
  }

  getProfile() {
    this.profileService.getAccount(this.store.getUsername()).subscribe(data => {
      console.log(data);
      this.profile = data;
    })
  }

  redirectToChangePass() {
    this.router.navigateByUrl("/changePass")
  }

  resend() {
    this.userService.resendCCodeAuth(this.store.getUsername()).subscribe(data => {
        alert("check you email")
        console.log(data)
    })
  }
  
}
