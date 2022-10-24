import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-tweet-popup',
  templateUrl: './tweet-popup.component.html',
  styleUrls: ['./tweet-popup.component.css']
})
export class TweetPopupComponent implements OnInit {

  constructor(private modalService: NgbModal) { }

  ngOnInit(): void {
  }

  openLg(content : any) {
    this.modalService.open(content, { size: 'lg' });
  }

  tweet : any = new Object;
  newTweet() {}
}
