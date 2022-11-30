import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environments/environment';
import { Tweet } from '../model/tweet';

@Injectable({
  providedIn: 'root'
})
export class TweetService {

  constructor(private client: HttpClient) { }

  getTweets() : Observable<Tweet[]> {
    return this.client.get<Tweet[]>(environment.apiUrl + "tweet_service/tweets");
  }
  
}
