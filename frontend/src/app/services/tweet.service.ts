import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError } from 'rxjs';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environments/environment';
import { Tweet } from '../model/tweet';
import { TweetLikes } from '../model/tweetLikes';

@Injectable({
  providedIn: 'root'
})
export class TweetService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getTweets() : Observable<Tweet[]> {
    return this.client.get<Tweet[]>(environment.apiUrl + "tweet_service/tweets");
  }

  postTweet(tweet : Tweet){
    console.log(tweet)
    return this.client.post<unknown>(environment.apiUrl + "tweet_service/tweets", {
      username: tweet.username,
      text: tweet.text,
    }, this.options())
  }

  likeTweet(tweet : Tweet) {
    console.log(tweet)
    return this.client.post<unknown>(environment.apiUrl + "tweet_service/tweets/likes", {
      tweetId: tweet.tweetId,
      username: tweet.username,
    }, this.options())
  }

  getLikes(tweet : Tweet) : Observable<TweetLikes>{
    console.log(tweet)
    return this.client.get<TweetLikes>(environment.apiUrl + "tweet_service/tweets/likes/" + tweet.tweetId)
  }
  
}