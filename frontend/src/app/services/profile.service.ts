import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environments/environment';
import { Profile } from '../model/profile';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getAccount(username: string): Observable<Profile> {
    return this.client.get<Profile>(environment.apiUrl + "profile_service/profile/" + username)
  }

  regUserRegular(username: string, password: string, email : string) {
    return this.client.post<unknown>(environment.apiUrl + "profile_service/regular", {
      username: username,
      password: password,
      email : email,
    }, this.options())
  }

  regUserBusiness(username: string, password: string, email : string) {
    return this.client.post<unknown>(environment.apiUrl + "profile_service/business", {
      username: username,
      password: password,
      email : email,
    }, this.options())
  }
}
