import { Component, Input, OnInit } from '@angular/core';
import { Tweet } from '../model/tweet';

@Component({
  selector: 'app-tweet',
  templateUrl: './tweet.component.html',
  styleUrls: ['./tweet.component.css']
})
export class TweetComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  @Input()
  tweet! : Tweet

}
