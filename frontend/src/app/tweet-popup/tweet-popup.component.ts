import { Component} from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Tweet } from '../model/tweet';
import { StoreService } from '../services/store-service.service';
import { TweetService } from '../services/tweet.service';

@Component({
  selector: 'app-tweet-popup',
  templateUrl: './tweet-popup.component.html',
  styleUrls: ['./tweet-popup.component.css']
})
export class TweetPopupComponent {

  constructor(private modalService: NgbModal, private store : StoreService, private tweetService : TweetService) { }

  openLg(content : any) {
    this.modalService.open(content, { size: 'lg' });
  }

  tweet : Tweet = new Tweet();

  newTweet() {
    this.tweet.username = this.store.getUsername();
    this.tweetService.postTweet(this.tweet).subscribe(data => {
      console.log(data);
    })
  }
}
